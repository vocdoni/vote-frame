package mongo

import (
	"fmt"
	"reflect"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

var (
	ErrUserUnknown     = fmt.Errorf("user unknown")
	ErrElectionUnknown = fmt.Errorf("electionID unknown")
)

// Users is the list of users.
type Users struct {
	Users []uint64 `json:"users"`
}

// User represents a farcaster user.
type User struct {
	UserID         uint64    `json:"userID,omitempty" bson:"_id"`
	ElectionCount  uint64    `json:"electionCount" bson:"electionCount"`
	CastedVotes    uint64    `json:"castedVotes" bson:"castedVotes"`
	Username       string    `json:"username" bson:"username"`
	Displayname    string    `json:"displayname" bson:"displayname"`
	CustodyAddress string    `json:"custodyAddress" bson:"custodyAddress"`
	Addresses      []string  `json:"addresses" bson:"addresses"`
	Signers        []string  `json:"signers" bson:"signers"`
	Followers      uint64    `json:"followers" bson:"followers"`
	LastUpdated    time.Time `json:"lastUpdated" bson:"lastUpdated"`
}

// UserAccessProfile holds the user's access profile data, used by our backend to determine the user's access level.
// It also holds the notification status.
type UserAccessProfile struct {
	UserID                  uint64   `json:"userID,omitempty" bson:"_id"`
	NotificationsAccepted   bool     `json:"notificationsAccepted" bson:"notificationsAccepted"`
	NotificationsRequested  bool     `json:"notificationsRequested" bson:"notificationsRequested"`
	Reputation              uint32   `json:"reputation" bson:"reputation"`
	AccessLevel             uint32   `json:"accessLevel" bson:"accessLevel"`
	WhiteListed             bool     `json:"whiteListed" bson:"whiteListed"`
	NotificationsMutedUsers []uint64 `json:"notificationsMutedUsers" bson:"notificationsMutedUsers"`
}

// Election represents an election and its details owned by a user.
type Election struct {
	ElectionMeta
	ElectionID            string    `json:"electionId" bson:"_id"`
	UserID                uint64    `json:"userId" bson:"userId"`
	CastedVotes           uint64    `json:"castedVotes" bson:"castedVotes"`
	LastVoteTime          time.Time `json:"lastVoteTime" bson:"lastVoteTime"`
	CreatedTime           time.Time `json:"createdTime" bson:"createdTime"`
	Source                string    `json:"source" bson:"source"`
	FarcasterUserCount    uint32    `json:"farcasterUserCount" bson:"farcasterUserCount"`
	InitialAddressesCount uint32    `json:"initialAddressesCount" bson:"initialAddressesCount"`
	Question              string    `json:"question" bson:"question"`
}

// Census stores the census of an election ready to be used for voting on farcaster.
type Census struct {
	CensusID           string            `json:"censusId" bson:"_id"`
	Root               string            `json:"root" bson:"root"`
	ElectionID         string            `json:"electionId" bson:"electionId"`
	TokenDecimals      uint32            `json:"tokenDecimals" bson:"tokenDecimals"`
	Participants       map[string]string `json:"participants" bson:"participants"`
	FromTotalAddresses uint32            `json:"fromTotalAddresses" bson:"fromTotalAddresses"`
	CreatedBy          uint64            `json:"createdBy" bson:"createdBy"`
}

// ElectionMeta stores non related election information that is useful
// for certain types of frame interactions
type ElectionMeta struct {
	// CensusERC20TokenDecimals is the number of decimals that a certain ERC20 token, that was used
	// for creating the census of the election, has.
	CensusERC20TokenDecimals uint32 `json:"censusERC20TokenDecimals" bson:"censusERC20TokenDecimals"`
}

// Results represents the final results of an election.
type Results struct {
	ElectionID string `json:"electionId" bson:"_id"`
	FinalPNG   []byte `json:"finalPNG" bson:"finalPNG"`
}

// VotersOfElection represents the list of voters of an election.
type VotersOfElection struct {
	ElectionID string   `json:"electionId" bson:"_id"`
	Voters     []uint64 `json:"voters" bson:"voters"`
}

// Authentication represents the authentication data for a user.
type Authentication struct {
	UserID     uint64    `json:"userId" bson:"_id"`
	AuthTokens []string  `json:"authTokens" bson:"authTokens"`
	UpdatedAt  time.Time `json:"updatedAt" bson:"updatedAt"`
}

// NotificationType represents the type of notification to be sent to a user.
type NotificationType int

const (
	NotificationTypeNewElection NotificationType = iota
	// create more notification types here
)

// Notification represents a notification to be sent to a user.
type Notification struct {
	ID             int64            `json:"id" bson:"_id"`
	Type           NotificationType `json:"type" bson:"type"`
	UserID         uint64           `json:"userId" bson:"userId"`
	Username       string           `json:"username" bson:"username"`
	AuthorID       uint64           `json:"authorId" bson:"authorId"`
	AuthorUsername string           `json:"authorUsername" bson:"authorUsername"`
	ElectionID     string           `json:"electionId" bson:"electionId"`
	FrameUrl       string           `json:"frameUrl" bson:"frameUrl"`
	CustomText     string           `json:"customText" bson:"customText"`
	Deadline       time.Time        `json:"deadline" bson:"deadline"`
}

// Collection is a dataset containing several users, elections and results (used for dump and import).
type Collection struct {
	UserCollection
	ElectionCollection
	ResultsCollection
	VotersOfElectionCollection
	CensusCollection
	CommunitiesCollection
}

// UserCollection is a dataset containing several users (used for dump and import).
type UserCollection struct {
	Users []User `json:"users" bson:"users"`
}

// ElectionCollection is a dataset containing several elections (used for dump and import).
type ElectionCollection struct {
	Elections []Election `json:"elections" bson:"elections"`
}

// CensusCollection is a dataset containing several censuses (used for dump and import).
type CensusCollection struct {
	Censuses []Census `json:"censuses" bson:"censuses"`
}

// ResultsCollection is a dataset containing several election results (used for dump and import).
type ResultsCollection struct {
	Results []Results `json:"results" bson:"results"`
}

// VotersOfElectionCollection is a dataset containing several voters of elections (used for dump and import).
type VotersOfElectionCollection struct {
	VotersOfElection []VotersOfElection `json:"votersOfElection" bson:"votersOfElection"`
}

// CommunitiesCollection is a dataset containing several communities (used for dump and import).
type CommunitiesCollection struct {
	Communities []Community `json:"communities" bson:"communities"`
}

type MetadataCollection struct {
	Metadata []Metadata `json:"metadata" bson:"metadata"`
}

// UserRanking is a user ranking entry.
type UserRanking struct {
	FID         uint64 `json:"fid" bson:"fid"`
	Username    string `json:"username" bson:"username"`
	Displayname string `json:"displayname" bson:"displayname"`
	Count       uint64 `json:"count" bson:"count"`
}

// ElectionRanking is an election ranking entry.
type ElectionRanking struct {
	ElectionID           string `json:"electionId" bson:"_id"`
	VoteCount            uint64 `json:"voteCount" bson:"voteCount"`
	CreatedByFID         uint64 `json:"createdByFID" bson:"createdByFID"`
	CreatedByUsername    string `json:"createdByUsername" bson:"createdByUsername"`
	CreatedByDisplayname string `json:"createdByDisplayname" bson:"createdByDisplayname"`
	Title                string `json:"title" bson:"title"`
}

// Community represents a community entry.
type Community struct {
	ID            uint64          `json:"id" bson:"_id"`
	Name          string          `json:"name" bson:"name"`
	Channels      []string        `json:"channels" bson:"channels"`
	Census        CommunityCensus `json:"census" bson:"census"`
	ImageURL      string          `json:"imageURL" bson:"imageURL"`
	Admins        []uint64        `json:"owners" bson:"owners"`
	Notifications bool            `json:"notifications" bson:"notifications"`
}

const (
	// TypeCommunityCensusChannel is the type for a community census that uses
	// a channel as source.
	TypeCommunityCensusChannel = "channel"
	// TypeCommunityCensusERC20 is the type for a community census that uses
	// ERC20 holders as source.
	TypeCommunityCensusERC20 = "erc20"
	// TypeCommunityCensusNFT is the type for a community census that uses
	// NFT holders as source.
	TypeCommunityCensusNFT = "nft"
)

type CommunityCensus struct {
	Name      string               `json:"name" bson:"name"`
	Type      string               `json:"type" bson:"type"`
	Addresses []CommunityAddresses `json:"addresses" bson:"addresses"`
	Channel   string               `json:"channel" bson:"channel"`
}

type CommunityAddresses struct {
	Address    string `json:"address" bson:"address"`
	Blockchain string `json:"blockchain" bson:"blockchain"`
}

type Metadata struct {
	Key   string `json:"key" bson:"_id"`
	Value any    `json:"value" bson:"value"`
}

// dynamicUpdateDocument creates a BSON update document from a struct, including only non-zero fields.
// It uses reflection to iterate over the struct fields and create the update document.
// The struct fields must have a bson tag to be included in the update document.
// The _id field is skipped.
func dynamicUpdateDocument(item interface{}, alwaysUpdateTags []string) (bson.M, error) {
	val := reflect.ValueOf(item)
	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}
	if !val.IsValid() || val.Kind() != reflect.Struct {
		return nil, fmt.Errorf("input must be a valid struct")
	}

	update := bson.M{}
	typ := val.Type()

	// Create a map for quick lookup
	alwaysUpdateMap := make(map[string]bool, len(alwaysUpdateTags))
	for _, tag := range alwaysUpdateTags {
		alwaysUpdateMap[tag] = true
	}

	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)
		if !field.CanInterface() {
			continue
		}
		fieldType := typ.Field(i)
		tag := fieldType.Tag.Get("bson")
		if tag == "" || tag == "-" {
			continue
		}

		// Check if the field should always be updated or is not the zero value
		_, alwaysUpdate := alwaysUpdateMap[tag]
		if alwaysUpdate || !reflect.DeepEqual(field.Interface(), reflect.Zero(field.Type()).Interface()) {
			update[tag] = field.Interface()
		}
	}

	return bson.M{"$set": update}, nil
}
