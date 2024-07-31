package reputation

import "github.com/ethereum/go-ethereum/common"

const (
	// user activity ponderation values
	followersDividerPonderation      = 2000
	electionsDividerPonderation      = 10
	votesDividerPonderation          = 4
	castedDividerPonderation         = 20
	communitiesMultiplierPonderation = 2
	// User activity max reputation values
	maxFollowersReputation = 10
	maxElectionsReputation = 10
	maxVotesReputation     = 25
	maxCastedReputation    = 45
	maxCommunityReputation = 10
	maxReputation          = 100
	// Boosters puntuaction values
	votecasterNFTPassPuntuaction              = 10
	votecasterLaunchNFTPuntuaction            = 18
	votecasterAlphafrensFollowerPuntuaction   = 12
	votecasterFarcasterFollowerPuntuaction    = 5
	vocdoniFarcasterFollowerPuntuaction       = 3
	votecasterAnnouncementRecastedPuntuaction = 7
	kiwiPuntuaction                           = 4
	degenDAONFTPuntuaction                    = 6
	haberdasheryNFTPuntuaction                = 6
	degenAtLeast10kPuntuaction                = 3
	tokyoDAONFTPuntuaction                    = 4
	proxyStudioNFTPuntuaction                 = 5
	proxyAtLeast5Puntuaction                  = 3
	nameDegenPuntuaction                      = 4
	farcasterOGNFTPuntuaction                 = 6
	moxiePassPuntuaction                      = 4

	// yield rate
	yieldParamA         = 2
	yieldParamB         = .2
	daoMultiplier       = 4
	channelMultiplier   = 2
	voterMultiplier     = .3
	ownerMultiplier     = .7
	communityMultiplier = 1
)

// ActivityPuntuationInfo contains the max reputation values for each activity
var ActivityPuntuationInfo = ReputationInfo{
	"maxFollowersReputation": maxFollowersReputation,
	"maxElectionsReputation": maxElectionsReputation,
	"maxVotesReputation":     maxVotesReputation,
	"maxCastedReputation":    maxCastedReputation,
	"maxCommunityReputation": maxCommunityReputation,
	"maxReputation":          maxReputation,
}

// BoostersPuntuationInfo contains the puntuaction values for each booster
var BoostersPuntuationInfo = ReputationInfo{
	"votecasterNFTPassPuntuaction":              votecasterNFTPassPuntuaction,
	"votecasterLaunchNFTPuntuaction":            votecasterLaunchNFTPuntuaction,
	"votecasterAlphafrensFollowerPuntuaction":   votecasterAlphafrensFollowerPuntuaction,
	"votecasterFarcasterFollowerPuntuaction":    votecasterFarcasterFollowerPuntuaction,
	"vocdoniFarcasterFollowerPuntuaction":       vocdoniFarcasterFollowerPuntuaction,
	"votecasterAnnouncementRecastedPuntuaction": votecasterAnnouncementRecastedPuntuaction,
	"kiwiPuntuaction":                           kiwiPuntuaction,
	"degenDAONFTPuntuaction":                    degenDAONFTPuntuaction,
	"haberdasheryNFTPuntuaction":                haberdasheryNFTPuntuaction,
	"degenAtLeast10kPuntuaction":                degenAtLeast10kPuntuaction,
	"tokyoDAONFTPuntuaction":                    tokyoDAONFTPuntuaction,
	"proxyStudioNFTPuntuaction":                 proxyStudioNFTPuntuaction,
	"proxyAtLeast5Puntuaction":                  proxyAtLeast5Puntuaction,
	"nameDegenPuntuaction":                      nameDegenPuntuaction,
	"farcasterOGNFTPuntuaction":                 farcasterOGNFTPuntuaction,
	"moxiePassPuntuaction":                      moxiePassPuntuaction,
}

// Boosters contract addresses
var (
	// Votecaster NFT Pass contract address
	// TODO: update
	VotecasterNFTPassAddress = common.HexToAddress("0x225D58E18218E8d87f365301aB6eEe4CbfAF820b")
	// Votecaster Launch NFT contract address
	// TODO: update
	VotecasterLaunchNFTAddress = common.HexToAddress("0x32B6BB4d1f7298d4a80c2Ece237e4474C0880B69")
	// Votecaster Alphafrens Channel address
	VotecasterAlphafrensChannelAddress = common.HexToAddress("0xa630fcc62165a3587c6857d73b556c8a61c8edd3")
	// $KIWI token contract address
	KIWIAddress = common.HexToAddress("0x66747bdC903d17C586fA09eE5D6b54CC85bBEA45")
	// DegenDAO NFT contract address
	DegenDAONFTAddress = common.HexToAddress("0x980Fbdd1cF05080781Dca0AEf7026B0406743389")
	// Haberdashery NFT contract address
	HaberdasheryNFTAddress = common.HexToAddress("0x85E7DF5708902bE39891d59aBEf8E21EDE91E8BF")
	// Degen token contract address
	DegenAddress = common.HexToAddress("0x4ed4E862860beD51a9570b96d89aF5E1B0Efefed")
	// TokyoDAO NFT contract address
	TokyoDAONFTAddress = common.HexToAddress("0x432073397Aead241cf2411e21D8fA949183E7151")
	// $PROXY token contract address
	ProxyAddress = common.HexToAddress("0xA051A2Cb19C00eCDffaE94D0Ff98c17758041D16")
	// ProxyStudio NFT contract address
	ProxyStudioNFTAddress = common.HexToAddress("0x7888b1f446c912ddec9bf582629e9ae8845fd8c6")
	// NameDegen NFT contract address
	NameDegenAddress = common.HexToAddress("0x4087fb91A1fBdef05761C02714335D232a2Bf3a1")
	// FarCaster OG NFT contract address
	FarcasterOGNFTAddress = common.HexToAddress("0xe03ef4b9db1a47464de84fb476f9baf493b3e886")
	// Moxie Pass NFT contract address:
	MoxiePassAddress = common.HexToAddress("0x235CAD50d8a510Bc9081279996f01877827142D8")
)

// Boosters costants (ids, hashesh and network information)
const (
	// Votecaster NFT Pass network short name and ID
	VotecasterNFTPassChainShortName = "base"
	VotecasterNFTPassChainID        = 8453
	// Votecaster Launch NFT network short name and ID
	VotecasterLaunchNFTChainShortName = "base"
	VotecasterLaunchNFTChainID        = 8453
	// Votecaster Farcaster ID
	VotecasterFarcasterFID uint64 = 521116
	// Vocdoni Farcaster ID
	VocdoniFarcasterFID uint64 = 7548
	// Votecaster Announcement Farcaster Cast Hash
	// TODO: update
	VotecasterAnnouncementCastHash = "0xe4528c4931127eb32e4c7c473622d4e3a1c6b0a3"
	// $KIWI token network ID
	KIWIChainID uint64 = 10
	// DegenDAO NFT network short name and ID
	DegenDAONFTChainShortName = "base"
	DegenDAONFTChainChainID   = 8453
	// Haberdashery NFT network short name and ID
	HaberdasheryNFTChainShortName = "base"
	HaberdasheryNFTChainChainID   = 8453
	// Degen token network short name and ID
	DegenChainShortName = "base"
	DegenChainID        = 8453
	// TokyoDAO NFT network short name and ID
	TokyoDAONFTChainShortName = "base"
	TokyoDAONFTChainChainID   = 8453
	// $PROXY token network short name and ID
	ProxyChainShortName = "degen"
	ProxyChainID        = 666666666
	// ProxyStudio NFT network short name and ID
	ProxyStudioNFTShortName = "base"
	ProxyStudioNFTChainID   = 8453
	// NameDegen NFT network short name and ID
	NameDegenChainShortName = "degen"
	NameDegenChainID        = 666666666
	// FarCaster OG NFT network short name and ID
	FarcasterOGNFTChainShortName = "zora"
	FarcasterOGNFTChainID        = 7777777
	// Moxie Pass NFT network short name and ID
	MoxiePassChainShortName = "base"
	MoxiePassChainChainID   = 8453
)
