package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net"
	"net/http"
	"os"
	"os/signal"
	"path"
	"strings"
	"syscall"
	"time"

	"github.com/google/uuid"
	flag "github.com/spf13/pflag"
	"github.com/spf13/viper"
	"github.com/vocdoni/vote-frame/discover"
	"github.com/vocdoni/vote-frame/farcasterapi"
	"github.com/vocdoni/vote-frame/farcasterapi/hub"
	"github.com/vocdoni/vote-frame/farcasterapi/neynar"
	"github.com/vocdoni/vote-frame/mongo"
	urlapi "go.vocdoni.io/dvote/api"
	"go.vocdoni.io/dvote/httprouter"
	"go.vocdoni.io/dvote/httprouter/apirest"
	"go.vocdoni.io/dvote/log"
)

var (
	serverURL   = "http://localhost:8888"
	explorerURL = "https://dev.explorer.vote"
	onvoteURL   = "https://dev.onvote.app"
)

func main() {
	flag.String("server", serverURL, "The full URL of the server (http or https)")
	flag.Bool("tlsDomain", false, "Should a TLS certificate be fetched from letsencrypt for the domain? (requires port 443)")
	flag.String("tlsDirCert", "", "The directory to use to store the TLS certificate")
	flag.String("listenHost", "0.0.0.0", "The host to listen on")
	flag.Int("listenPort", 8888, "The port to listen on")
	flag.String("dataDir", "", "The directory to use for the data")
	flag.String("apiEndpoint", "https://api-dev.vocdoni.net/v2", "The Vocdoni API endpoint to use")
	flag.String("vocdoniPrivKey", "", "The Vocdoni private key to use for orchestrating the election (hex)")
	flag.String("censusFromFile", "farcaster_census.json", "Take census details from JSON file")
	flag.String("logLevel", "info", "The log level to use")
	flag.String("web", "./webapp/dist", "The path where the static web app is located")
	flag.String("explorerURL", explorerURL, "The full URL of the explorer (http or https)")
	flag.String("onvoteURL", onvoteURL, "The full URL of the onvote.app application (http or https)")
	flag.String("mongoURL", "", "The URL of the MongoDB server")
	flag.String("mongoDB", "voteframe", "The name of the MongoDB database")
	flag.String("adminToken", "", "The admin token to use for the API (if not set, it will be generated)")
	flag.Int("pollSize", 0, "The maximum votes allowed per poll (the more votes, the more expensive) (0 for default)")
	flag.Int("pprofPort", 0, "The port to use for the pprof http endpoints")
	flag.String("web3", "https://mainnet.optimism.io", "Web3 RPC Optimism endpoint")

	// bot flags
	// DISCLAMER: Currently the bot needs a HUB with write permissions to work.
	// It also needs a FID to impersonate to it and its private key to sign the
	// casts. Alternatively, it can be used with a Neynar API, but due the last
	// issues with the Neynar API, it is not recommended.
	flag.Uint64("botFid", 0, "FID to be used for the bot")
	flag.String("botPrivKey", "", "The bot private key to use for signing the vote (hex)")
	flag.String("botHubEndpoint", "", "The hub endpoint to use")
	flag.String("neynarAPIKey", "", "neynar API key")
	flag.String("neynarSignerUUID", "", "neynar signer UUID")
	flag.String("neynarWebhookSecret", "", "neynar Webhook shared secret")

	// Parse the command line flags
	flag.Parse()

	// Initialize Viper
	viper.SetEnvPrefix("VOCDONI")
	if err := viper.BindPFlags(flag.CommandLine); err != nil {
		panic(err)
	}
	viper.AutomaticEnv()

	// Using Viper to access the variables
	server := viper.GetString("server")
	tlsDomain := viper.GetBool("tlsDomain")
	tlsDirCert := viper.GetString("tlsDirCert")
	host := viper.GetString("listenHost")
	port := viper.GetInt("listenPort")
	dataDir := viper.GetString("dataDir")
	apiEndpoint := viper.GetString("apiEndpoint")
	vocdoniPrivKey := viper.GetString("vocdoniPrivKey")
	censusFromFile := viper.GetString("censusFromFile")
	logLevel := viper.GetString("logLevel")
	webAppDir := viper.GetString("web")
	explorerURL = viper.GetString("explorerURL")
	onvoteURL = viper.GetString("onvoteURL")
	mongoURL := viper.GetString("mongoURL")
	mongoDB := viper.GetString("mongoDB")
	adminToken := viper.GetString("adminToken")
	pollSize := viper.GetInt("pollSize")
	pprofPort := viper.GetInt("pprofPort")
	web3endpoint := viper.GetString("web3")
	// bot vars
	botFid := viper.GetUint64("botFid")
	botPrivKey := viper.GetString("botPrivKey")
	botHubEndpoint := viper.GetString("botHubEndpoint")
	neynarAPIKey := viper.GetString("neynarAPIKey")
	neynarSignerUUID := viper.GetString("neynarSignerUUID")
	neynarWebhookSecret := viper.GetString("neynarWebhookSecret")

	if adminToken == "" {
		adminToken = uuid.New().String()
		fmt.Printf("generated admin token: %s\n", adminToken)
	}

	log.Init(logLevel, "stdout", nil)

	log.Infow("configuration loaded",
		"server", server,
		"tlsDomain", tlsDomain,
		"tlsDirCert", tlsDirCert,
		"host", host,
		"port", port,
		"dataDir", dataDir,
		"apiEndpoint", apiEndpoint,
		"censusFromFile", censusFromFile,
		"logLevel", logLevel,
		"webAppDir", webAppDir,
		"explorerURL", explorerURL,
		"onvoteURL", onvoteURL,
		"mongoURL", mongoURL,
		"mongoDB", mongoDB,
		"pollSize", pollSize,
		"pprofPort", pprofPort,
		"botFid", botFid,
		"botPrivKey", botPrivKey,
		"botHubEndpoint", botHubEndpoint,
		"neynarAPIKey", neynarAPIKey,
		"neynarSignerUUID", neynarSignerUUID,
		"neynarWebhookSecret", neynarWebhookSecret,
		"web3endpoint", web3endpoint,
	)

	// Start the pprof http endpoints
	if pprofPort > 0 {
		ln, err := net.Listen("tcp", fmt.Sprintf("0.0.0.0:%d", pprofPort))
		if err != nil {
			log.Fatal(err)
		}
		go func() {
			log.Warnf("started pprof http endpoints at http://%s/debug/pprof", ln.Addr())
			log.Error(http.Serve(ln, nil))
		}()
	}

	// check the server URL is http or https and extract the domain
	if !strings.HasPrefix(server, "http://") && !strings.HasPrefix(server, "https://") {
		log.Fatal("server URL must start with http:// or https://")
	}
	serverURL = server
	domain := strings.Split(serverURL, "/")[2]
	log.Infow("server URL", "URL", serverURL, "domain", domain)

	// Set the maximum election size based on the API endpoint (try to guess the environment)
	if pollSize > 0 {
		maxElectionSize = pollSize
	} else {
		if strings.Contains(apiEndpoint, "stg") {
			maxElectionSize = stageMaxElectionSize
		} else {
			if strings.Contains(apiEndpoint, "dev") {
				maxElectionSize = devMaxElectionSize
			}
		}
	}

	// Create or load the census
	censusInfo := &CensusInfo{}
	if censusFromFile != "" {
		if err := censusInfo.FromFile(censusFromFile); err != nil {
			log.Fatal(err)
		}
	} else {
		log.Fatal("censusFromFile is required")
	}
	if censusInfo.Url == "" || len(censusInfo.Root) == 0 || censusInfo.Size == 0 {
		log.Fatal("censusFromFile must contain a valid URL and root hash")
	}

	// Create the MongoDB connection
	db, err := mongo.New(mongoURL, mongoDB)
	if err != nil {
		log.Fatal(err)
	}

	// Create the Farcaster API client
	neynarcli, err := neynar.NewNeynarAPI(neynarAPIKey, web3endpoint)
	if err != nil {
		log.Fatal(err)
	}

	// Start the discovery user profile background process
	mainCtx, mainCtxCancel := context.WithCancel(context.Background())
	discover.NewFarcasterDiscover(db, neynarcli).Run(mainCtx)
	defer mainCtxCancel()

	// Create the Vocdoni handler
	handler, err := NewVocdoniHandler(apiEndpoint, vocdoniPrivKey, censusInfo, webAppDir, db, mainCtx, neynarcli)
	if err != nil {
		log.Fatal(err)
	}

	// Create the HTTP API router
	router := new(httprouter.HTTProuter)
	if tlsDomain {
		router.TLSdomain = domain
	}

	router.TLSdirCert = tlsDirCert
	if err := router.Init(host, port); err != nil {
		log.Fatal(err)
	}

	// Add handler to serve the static files
	log.Infow("serving webapp static files from", "dir", webAppDir)
	// check index.html exists
	if _, err := os.Stat(path.Join(webAppDir, "index.html")); err != nil {
		log.Fatalf("index.html not found in webapp directory %s", webAppDir)
	}
	router.AddRawHTTPHandler("/app*", http.MethodGet, handler.staticHandler)
	router.AddRawHTTPHandler("/favicon.ico", http.MethodGet, func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, path.Join(webAppDir, "favicon.ico"))
	})

	// Add the Prometheus endpoint
	router.ExposePrometheusEndpoint("/metrics")

	// Create the API handler
	uAPI, err := urlapi.NewAPI(router, "/", dataDir, "")
	if err != nil {
		log.Fatal(err)
	}
	// Set the admin token
	uAPI.Endpoint.SetAdminToken(adminToken)

	// The root endpoint redirects to /app
	if err := uAPI.Endpoint.RegisterMethod("/", http.MethodGet, "public", func(msg *apirest.APIdata, ctx *httprouter.HTTPContext) error {
		ctx.Writer.Header().Add("Location", "/app")
		return ctx.Send([]byte("Redirecting to /app"), http.StatusTemporaryRedirect)
	}); err != nil {
		log.Fatal(err)
	}

	if err := uAPI.Endpoint.RegisterMethod("/", http.MethodPost, "public", func(msg *apirest.APIdata, ctx *httprouter.HTTPContext) error {
		ctx.Writer.Header().Add("Location", "/app")
		return ctx.Send([]byte("Redirecting to /app"), http.StatusTemporaryRedirect)
	}); err != nil {
		log.Fatal(err)
	}

	if err := uAPI.Endpoint.RegisterMethod("/{electionID}", http.MethodPost, "public", handler.landing); err != nil {
		log.Fatal(err)
	}

	if err := uAPI.Endpoint.RegisterMethod("/{electionID}", http.MethodGet, "public", handler.landing); err != nil {
		log.Fatal(err)
	}

	if err := uAPI.Endpoint.RegisterMethod("/dumpdb", http.MethodGet, "admin", handler.dumpDB); err != nil {
		log.Fatal(err)
	}

	if err := uAPI.Endpoint.RegisterMethod("/importdb", http.MethodPost, "admin", handler.importDB); err != nil {
		log.Fatal(err)
	}

	if err := uAPI.Endpoint.RegisterMethod("/rankings/usersByCreatedPolls", http.MethodGet, "public", handler.rankingByElectionsCreated); err != nil {
		log.Fatal(err)
	}

	if err := uAPI.Endpoint.RegisterMethod("/rankings/usersByCastedVotes", http.MethodGet, "public", handler.rankingByVotes); err != nil {
		log.Fatal(err)
	}

	if err := uAPI.Endpoint.RegisterMethod("/rankings/pollsByVotes", http.MethodGet, "public", handler.rankingOfElections); err != nil {
		log.Fatal(err)
	}

	// Register the API methods
	if err := uAPI.Endpoint.RegisterMethod("/router/{electionID}", http.MethodPost, "public", func(msg *apirest.APIdata, ctx *httprouter.HTTPContext) error {
		electionID := ctx.URLParam("electionID")
		packet := &FrameSignaturePacket{}
		if err := json.Unmarshal(msg.Data, packet); err != nil {
			return fmt.Errorf("failed to unmarshal frame signature packet: %w", err)
		}
		redirectURL := ""
		switch packet.UntrustedData.ButtonIndex {
		case 1:
			redirectURL = fmt.Sprintf(serverURL+"/poll/results/%s", electionID)
		case 2:
			redirectURL = fmt.Sprintf(serverURL+"/poll/%s", electionID)
		case 3:
			redirectURL = fmt.Sprintf(serverURL+"/info/%s", electionID)
		default:
			redirectURL = serverURL + "/"
		}
		log.Infow("received router request", "electionID", electionID, "buttonIndex", packet.UntrustedData.ButtonIndex, "redirectURL", redirectURL)
		ctx.Writer.Header().Add("Location", redirectURL)
		return ctx.Send([]byte(redirectURL), http.StatusTemporaryRedirect)
	}); err != nil {
		log.Fatal(err)
	}

	if err := uAPI.Endpoint.RegisterMethod("/poll/results/{electionID}", http.MethodGet, "public", handler.results); err != nil {
		log.Fatal(err)
	}

	if err := uAPI.Endpoint.RegisterMethod("/poll/results/{electionID}", http.MethodPost, "public", handler.results); err != nil {
		log.Fatal(err)
	}

	if err := uAPI.Endpoint.RegisterMethod("/poll/{electionID}", http.MethodGet, "public", handler.showElection); err != nil {
		log.Fatal(err)
	}

	if err := uAPI.Endpoint.RegisterMethod("/poll/{electionID}", http.MethodPost, "public", handler.showElection); err != nil {
		log.Fatal(err)
	}

	if err := uAPI.Endpoint.RegisterMethod("/vote/{electionID}", http.MethodPost, "public", handler.vote); err != nil {
		log.Fatal(err)
	}

	if err := uAPI.Endpoint.RegisterMethod("/vote/{electionID}", http.MethodGet, "public", handler.vote); err != nil {
		log.Fatal(err)
	}

	if err := uAPI.Endpoint.RegisterMethod("/info/{electionID}", http.MethodGet, "public", handler.info); err != nil {
		log.Fatal(err)
	}

	if err := uAPI.Endpoint.RegisterMethod("/info/{electionID}", http.MethodPost, "public", handler.info); err != nil {
		log.Fatal(err)
	}

	if err := uAPI.Endpoint.RegisterMethod("/create", http.MethodPost, "public", handler.createElection); err != nil {
		log.Fatal(err)
	}

	if err := uAPI.Endpoint.RegisterMethod("/census/csv", http.MethodPost, "public", handler.censusCSV); err != nil {
		log.Fatal(err)
	}

	if err := uAPI.Endpoint.RegisterMethod("/census/check/{censusID}", http.MethodGet, "public", handler.censusQueueInfo); err != nil {
		log.Fatal(err)
	}

	if err := uAPI.Endpoint.RegisterMethod("/create/check/{electionID}", http.MethodGet, "public", handler.checkElection); err != nil {
		log.Fatal(err)
	}

	if err := uAPI.Endpoint.RegisterMethod("/votersOf/{electionID}", http.MethodGet, "public", handler.votersForElection); err != nil {
		log.Fatal(err)
	}

	if err := uAPI.Endpoint.RegisterMethod("/testimage", http.MethodGet, "public", handler.testImage); err != nil {
		log.Fatal(err)
	}

	if err := uAPI.Endpoint.RegisterMethod("/testimage", http.MethodPost, "public", handler.testImage); err != nil {
		log.Fatal(err)
	}

	if err := uAPI.Endpoint.RegisterMethod("/preview/{electionID}", http.MethodGet, "public", handler.preview); err != nil {
		log.Fatal(err)
	}

	if err := uAPI.Endpoint.RegisterMethod(fmt.Sprintf("%s/{id}.png", imageHandlerPath), http.MethodGet, "public", handler.imagesHandler); err != nil {
		log.Fatal(err)
	}
	// if a bot FID is provided, start the bot background process
	if botFid > 0 {
		var botAPI farcasterapi.API
		if botPrivKey != "" && botHubEndpoint != "" {
			// Hub based bot
			botAPI, err = hub.NewHubAPI(botHubEndpoint, nil)
			if err != nil {
				log.Fatal(err)
			}
			if err := botAPI.SetFarcasterUser(botFid, botPrivKey); err != nil {
				log.Fatal(err)
			}
			log.Info("trying to init Hub based bot")
		} else if neynarAPIKey != "" && neynarSignerUUID != "" && neynarWebhookSecret != "" {
			// Neynar based bot
			if err := neynarcli.SetFarcasterUser(botFid, neynarSignerUUID); err != nil {
				log.Fatal(err)
			}
			botAPI = neynarcli
			// register neynar webhook handler
			if err := uAPI.Endpoint.RegisterMethod("/webhook/neynar", http.MethodPost, "public", neynarWebhook(neynarcli, neynarWebhookSecret)); err != nil {
				log.Fatal(err)
			}
			log.Info("trying to init Neynar based bot")
		} else {
			log.Fatalf("botFid is set but botPrivKey and botHubEndpoint or neynarAPIKey, neynarSignerUUID and neynarWebhookSecret are not")
		}
		voteBot, err := initBot(mainCtx, handler, botAPI, censusInfo)
		if err != nil {
			log.Fatal(err)
		}
		defer voteBot.Stop()
		log.Info("bot started")
	}

	// close if interrupt received
	log.Infof("startup complete at %s", time.Now().Format(time.RFC850))
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	<-c
	log.Warnf("received SIGTERM, exiting at %s", time.Now().Format(time.RFC850))
}
