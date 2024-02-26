package neynar

import (
	"bytes"
	"context"
	"crypto/hmac"
	"crypto/sha512"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"sync"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/vocdoni/farcaster-poc/farcasterapi"
	"go.vocdoni.io/dvote/log"
	"go.vocdoni.io/dvote/util"
)

const (
	neynarHubEndpoint = "https://hub-api.neynar.com"
	neynarAPIEndpoint = "https://api.neynar.com"

	// endpoints
	neynarGetUsernameEndpoint = neynarAPIEndpoint + "/v1/farcaster/user?fid=%d"
	neynarGetCastsEndpoint    = neynarAPIEndpoint + "/v1/farcaster/mentions-and-replies?fid=%d&limit=150&cursor=%s"
	neynarReplyEndpoint       = neynarAPIEndpoint + "/v2/farcaster/cast"
	neynarUserByEthAddresses  = neynarAPIEndpoint + "/v2/farcaster/user/bulk-by-address?addresses=%s"
	neynarVerificationsByFID  = neynarHubEndpoint + "/v1/verificationsByFid?fid=%d"

	MaxAddressesPerRequest = 300

	// timeouts
	getBotUsernameTimeout   = 10 * time.Second
	getCastByMentionTimeout = 60 * time.Second
	postCastTimeout         = 10 * time.Second

	// Requests backoff parameters
	maxConcurrentRequests = 2
	maxRetries            = 12              // Maximum number of retries
	baseDelay             = 1 * time.Second // Initial delay, increases exponentially

	// other
	neynarMentionType     = "cast-mention"
	neynarCastCreatedType = "cast.created"
	neynarCastType        = "cast"
	timeLayout            = "2006-01-02T15:04:05.000Z"
)

type NeynarAPI struct {
	fid          uint64
	username     string
	signerUUID   string
	apiKey       string
	reqSemaphore chan struct{} // Semaphore to limit concurrent requests
	newCasts     map[uint64]*farcasterapi.APIMessage
	newCastsMtx  sync.Mutex
}

func NewNeynarAPI(apiKey string) *NeynarAPI {
	return &NeynarAPI{
		apiKey:       apiKey,
		reqSemaphore: make(chan struct{}, maxConcurrentRequests),
		newCasts:     make(map[uint64]*farcasterapi.APIMessage),
	}
}

// SetFarcasterUser method sets the farcaster user with the given fid and signer.
// The signer is the UUID of the user that signs the messages.
func (n *NeynarAPI) SetFarcasterUser(fid uint64, signer string) error {
	n.fid = fid
	n.signerUUID = signer
	var err error
	ctx, cancel := context.WithTimeout(context.Background(), getBotUsernameTimeout)
	defer cancel()
	userdata, err := n.UserDataByFID(ctx, n.fid)
	if err != nil {
		return fmt.Errorf("error getting bot username: %w", err)
	}
	n.username = userdata.Username
	return nil
}

func (n *NeynarAPI) Stop() error {
	return nil
}

func (n *NeynarAPI) LastMentions(ctx context.Context, timestamp uint64) ([]*farcasterapi.APIMessage, uint64, error) {
	if n.fid == 0 {
		return nil, 0, fmt.Errorf("farcaster user not set")
	}
	// get new mentions from the queue and calculate the last timestamp
	messages := []*farcasterapi.APIMessage{}
	lastTimestamp := timestamp
	n.newCastsMtx.Lock()
	for ts, msg := range n.newCasts {
		if ts > timestamp {
			messages = append(messages, msg)
			lastTimestamp = ts
		}
	}
	// clear the new mentions queue
	n.newCasts = make(map[uint64]*farcasterapi.APIMessage)
	n.newCastsMtx.Unlock()
	return messages, lastTimestamp, nil
}

func (n *NeynarAPI) Reply(ctx context.Context, fid uint64, parentHash, content string) error {
	if n.fid == 0 {
		return fmt.Errorf("farcaster user not set")
	}
	// create request body
	castReq := &CastPostRequest{
		Signer: n.signerUUID,
		Text:   content,
		Parent: parentHash,
	}
	body, err := json.Marshal(castReq)
	if err != nil {
		return fmt.Errorf("error marshalling request body: %w", err)
	}
	// create request with the bot fid and set the api key header
	_, err = n.request(neynarReplyEndpoint, http.MethodPost, body)
	return err
}

// UserData method returns the username, the custody address and the
// verification addresses of the user with the given fid.
func (n *NeynarAPI) UserDataByFID(ctx context.Context, fid uint64) (*farcasterapi.Userdata, error) {
	// create request with the bot fid
	url := fmt.Sprintf(neynarGetUsernameEndpoint, fid)
	body, err := n.request(url, http.MethodGet, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}
	// decode username
	usernameResponse := &UserdataResponse{}
	if err := json.Unmarshal(body, usernameResponse); err != nil {
		return nil, fmt.Errorf("error unmarshalling response body: %w", err)
	}

	// get signers
	signers, err := n.signersFromFid(fid)
	if err != nil {
		return nil, fmt.Errorf("error getting signers: %w", err)
	}

	return &farcasterapi.Userdata{
		FID:                    fid,
		Username:               usernameResponse.Result.User.Username,
		CustodyAddress:         usernameResponse.Result.User.CustodyAddress,
		VerificationsAddresses: usernameResponse.Result.User.VerificationsAddresses,
		Signers:                signers,
	}, nil
}

func (n *NeynarAPI) signersFromFid(fid uint64) ([]string, error) {
	// get verifications to fetch the signer
	body, err := n.request(fmt.Sprintf(neynarVerificationsByFID, fid), http.MethodGet, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}

	// decode verifications json
	verificationsData := &HubAPIResponse{}
	if err := json.Unmarshal(body, verificationsData); err != nil {
		return nil, fmt.Errorf("error unmarshalling verifications: %w", err)
	}
	// filter verifications addresses
	signersMap := make(map[string]struct{})
	for _, msg := range verificationsData.Messages {
		// if no data or verification data is found, skip. If the message data
		// type is not the one we are looking for, skip
		if msg.Data == nil || msg.Data.Type != HUB_MESSAGE_TYPE_VERIFICATION || msg.Data.VerificationAddEthAddressBody == nil || msg.Signer == "" {
			log.Warnw("invalid verification message", "msg", msg)
			continue
		}
		signersMap[msg.Signer] = struct{}{}
	}
	signers := []string{}
	for signer := range signersMap {
		signers = append(signers, strings.ToLower(signer))
	}
	return signers, nil
}

func (n *NeynarAPI) UserDataByVerificationAddress(ctx context.Context, addresses []string) ([]*farcasterapi.Userdata, error) {
	if len(addresses) > MaxAddressesPerRequest {
		return nil, fmt.Errorf("address slice exceeds the maximum limit of 350 addresses")
	}

	// Concatenate addresses separated by commas
	addressesStr := strings.Join(addresses, ",")

	// Construct the URL with multiple addresses
	url := fmt.Sprintf(neynarUserByEthAddresses, addressesStr)

	// Make the request
	body, err := n.request(url, http.MethodGet, nil)
	if err != nil {
		return nil, err
	}

	// Decode the response
	var results map[string][]UserdataV2
	if err := json.Unmarshal(body, &results); err != nil {
		return nil, fmt.Errorf("error unmarshalling response body: %w", err)
	}

	// Process results into []*farcasterapi.Userdata
	userDataSlice := make([]*farcasterapi.Userdata, 0)
	for address, dataItems := range results {
		for _, item := range dataItems {
			if item.Username != "" {
				if len(item.VerifiedAddresses.EthAddresses) == 0 {
					log.Warnw("no verified addresses found", "user", item.Username)
					continue
				}
				// Normalize addresses to the Ethereum hex standard format
				var normalizedAddresses []string
				for _, addr := range item.VerifiedAddresses.EthAddresses {
					normalizedAddresses = append(normalizedAddresses, common.HexToAddress(addr).Hex())
				}

				// Get signers
				signers, err := n.signersFromFid(item.Fid)
				if err != nil {
					return nil, fmt.Errorf("error getting signers for address %s: %w", address, err)
				}
				if len(signers) == 0 {
					log.Warnw("no signers found", "user", item.Username, "address", address)
					continue
				}
				userData := &farcasterapi.Userdata{
					FID:                    item.Fid,
					Username:               item.Username,
					CustodyAddress:         common.HexToAddress(item.CustodyAddress).Hex(),
					VerificationsAddresses: normalizedAddresses,
					Signers:                signers,
				}

				userDataSlice = append(userDataSlice, userData)
				break // we only need the first valid user data per address
			}
		}
	}

	if len(userDataSlice) == 0 {
		return nil, farcasterapi.ErrNoDataFound
	}

	return userDataSlice, nil
}

func (n *NeynarAPI) WebhookHandler(body []byte) error {
	// decode the request body
	castWebhookReq := &CastWebhookRequest{}
	if err := json.Unmarshal(body, castWebhookReq); err != nil {
		return fmt.Errorf("error unmarshalling request body: %s", err.Error())
	}
	// check if the req type is a not created cast or data type is not cast and
	// skip if so
	if castWebhookReq.Type != neynarCastCreatedType || castWebhookReq.Data.Object != neynarCastType {
		return nil
	}
	// parse timestamp
	parsedTimestamp, err := time.Parse(timeLayout, castWebhookReq.Data.Timestamp)
	if err != nil {
		return fmt.Errorf("error parsing timestamp: %s", err.Error())
	}
	notificationTimestamp := uint64(parsedTimestamp.Unix())
	// check if the cast is a mention and skip if not
	mention := fmt.Sprintf("@%s", n.username)
	if !strings.HasPrefix(castWebhookReq.Data.Text, mention) {
		return nil
	}
	// parse the text to remove the bot username and add mention to the
	// list
	text := strings.TrimSpace(strings.TrimPrefix(castWebhookReq.Data.Text, mention))
	message := &farcasterapi.APIMessage{
		IsMention: true,
		Author:    castWebhookReq.Data.Author.Fid,
		Content:   text,
		Hash:      castWebhookReq.Data.Hash,
	}
	// add the new mention to the list
	n.newCastsMtx.Lock()
	n.newCasts[notificationTimestamp] = message
	n.newCastsMtx.Unlock()
	return nil
}

func (n *NeynarAPI) request(url, method string, body []byte) ([]byte, error) {
	for attempt := 0; attempt < maxRetries; attempt++ {
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		req, err := http.NewRequestWithContext(ctx, method, url, bytes.NewReader(body))
		if err != nil {
			return nil, fmt.Errorf("error creating request: %w", err)
		}
		req.Header.Set("api_key", n.apiKey)
		req.Header.Set("Content-Type", "application/json")

		// We need to avoid too much concurrent requests and penalization from the API
		n.reqSemaphore <- struct{}{}
		res, err := http.DefaultClient.Do(req)
		<-n.reqSemaphore
		if err != nil {
			return nil, fmt.Errorf("error downloading json: %w", err)
		}
		defer res.Body.Close()
		if res.StatusCode == http.StatusTooManyRequests {
			time.Sleep(time.Duration(attempt+1)*baseDelay + time.Duration(util.RandomInt(0, 2000))*time.Millisecond)
		} else if res.StatusCode != http.StatusOK {
			return nil, fmt.Errorf("error downloading json: %s", res.Status)
		} else {
			respBody, err := io.ReadAll(res.Body)
			if err != nil {
				return nil, fmt.Errorf("error reading response body: %w", err)
			}
			return respBody, nil // Success
		}
		log.Debugw("retrying request", "attempt", attempt+1, "url", url, "method", method)
	}

	return nil, fmt.Errorf("error downloading json: exceeded retry limit")
}

// VerifyRequest method verifies the request signature and returns the body of
// the request, a boolean indicating if the signature is valid and an error if
// something goes wrong.
func VerifyRequest(secret, neynarSig string, body []byte) ([]byte, bool, error) {
	// Create HMAC with SHA512 and update it with the body
	hmac := hmac.New(sha512.New, []byte(secret))
	hmac.Write(body)
	// Calculate the HMAC signature and encode it in hexadecimal
	signature := hex.EncodeToString(hmac.Sum(nil))
	log.Debugw("neynar request", "body", string(body), "neynarSig", neynarSig, "signature", signature)
	// verify the signature
	return body, signature == neynarSig, nil
}
