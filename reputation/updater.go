package reputation

import (
	"context"
	"errors"
	"fmt"
	"math/big"
	"sync"
	"sync/atomic"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/vocdoni/census3/apiclient"
	"github.com/vocdoni/vote-frame/airstack"
	"github.com/vocdoni/vote-frame/alfafrens"
	"github.com/vocdoni/vote-frame/communityhub"
	"github.com/vocdoni/vote-frame/farcasterapi"
	"github.com/vocdoni/vote-frame/mongo"
	"go.vocdoni.io/dvote/log"
)

// Updater is a struct to update user reputation data in the database
// periodically. It calculates the reputation of each user based on their
// activity and boosters. It gets the activity data from the database and the
// boosters data from the Airstack and the Census3 API.
type Updater struct {
	ctx    context.Context
	cancel context.CancelFunc
	waiter sync.WaitGroup

	db            *mongo.MongoStorage
	fapi          farcasterapi.API
	airstack      *airstack.Airstack
	census3       *apiclient.HTTPclient
	lastUpdate    time.Time
	maxConcurrent int

	alfafrensFollowers  map[uint64]bool
	vocdoniFollowers    map[uint64]bool
	votecasterFollowers map[uint64]bool
	recasters           map[uint64]bool
	followersMtx        sync.Mutex

	votecasterNFTPassHolders   map[common.Address]*big.Int
	votecasterLaunchNFTHolders map[common.Address]*big.Int
	kiwiHolders                map[common.Address]*big.Int
	degenDAONFTHolders         map[common.Address]*big.Int
	haberdasheryNFTHolders     map[common.Address]*big.Int
	tokyoDAONFTHolders         map[common.Address]*big.Int
	proxyHolders               map[common.Address]*big.Int
	proxyStudioNFTHolders      map[common.Address]*big.Int
	nameDegenHolders           map[common.Address]*big.Int
	holdersMtx                 sync.Mutex
}

// NewUpdater creates a new Updater instance with the given parameters,
// including the parent context, the database, the Airstack client, the Census3
// client, and the maximum number of concurrent updates.
func NewUpdater(ctx context.Context, db *mongo.MongoStorage, fapi farcasterapi.API,
	as *airstack.Airstack, c3 *apiclient.HTTPclient, maxConcurrent int,
) (*Updater, error) {
	if db == nil {
		return nil, errors.New("database is required")
	}
	if fapi == nil {
		return nil, errors.New("farcaster api is required")
	}
	if as == nil {
		return nil, errors.New("airstack client is required")
	}
	if c3 == nil {
		return nil, errors.New("census3 client is required")
	}
	internalCtx, cancel := context.WithCancel(ctx)
	return &Updater{
		ctx:                        internalCtx,
		cancel:                     cancel,
		db:                         db,
		fapi:                       fapi,
		airstack:                   as,
		census3:                    c3,
		lastUpdate:                 time.Time{},
		maxConcurrent:              maxConcurrent,
		alfafrensFollowers:         make(map[uint64]bool),
		vocdoniFollowers:           make(map[uint64]bool),
		votecasterFollowers:        make(map[uint64]bool),
		recasters:                  make(map[uint64]bool),
		votecasterNFTPassHolders:   make(map[common.Address]*big.Int),
		votecasterLaunchNFTHolders: make(map[common.Address]*big.Int),
		kiwiHolders:                make(map[common.Address]*big.Int),
		degenDAONFTHolders:         make(map[common.Address]*big.Int),
		haberdasheryNFTHolders:     make(map[common.Address]*big.Int),
		tokyoDAONFTHolders:         make(map[common.Address]*big.Int),
		proxyHolders:               make(map[common.Address]*big.Int),
		proxyStudioNFTHolders:      make(map[common.Address]*big.Int),
		nameDegenHolders:           make(map[common.Address]*big.Int),
	}, nil
}

// Start method starts the updater with the given cooldown time between updates.
// It will run until the context is canceled, calling the updateUsers method
// periodically and updating the last update time accordingly.
func (u *Updater) Start(coolDown time.Duration) error {
	u.waiter.Add(1)
	go func() {
		defer u.waiter.Done()

		for {
			select {
			case <-u.ctx.Done():
				return
			default:
				// check if is time to update users
				if time.Since(u.lastUpdate) < coolDown {
					time.Sleep(time.Second * 30)
					continue
				}
				// update internal followers
				if err := u.updateFollowersAndRecasters(); err != nil {
					log.Warnw("error updating internal followers", "error", err)
				}
				// update holders
				if err := u.updateHolders(); err != nil {
					log.Warnw("error updating holders", "error", err)
				}
				// launch update communities
				if err := u.updateCommunities(); err != nil {
					log.Warnw("error updating communities", "error", err)
				}
				// launch update
				if err := u.updateUsers(); err != nil {
					log.Warnw("error updating users", "error", err)
				}
				// update last update time
				u.lastUpdate = time.Now()
			}
		}
	}()
	return nil
}

// Stop method stops the updater by canceling the context and waiting for the
// updater to finish.
func (u *Updater) Stop() {
	u.cancel()
	u.waiter.Wait()
}

// updateFollowersAndRecasters method updates the internal followers of the
// Vocdoni and Votecaster profiles in Farcaster and warpcast users that have
// recasted the Votecaster Launch cast announcement. It fetches the followers
// and recasters data from the Farcaster API and updates the internal followers
// maps accordingly. It returns an error if the followers data cannot be fetched.
func (u *Updater) updateFollowersAndRecasters() error {
	internalCtx, cancel := context.WithTimeout(u.ctx, time.Second*30)
	defer cancel()
	u.followersMtx.Lock()
	defer u.followersMtx.Unlock()
	// update alfafrens followers
	alfafrensFollowers, err := alfafrens.ChannelFids(VotecasterAlphafrensChannelAddress.Bytes())
	if err == nil {
		for _, fid := range alfafrensFollowers {
			u.alfafrensFollowers[fid] = true
		}
	}
	// update vocdoni followers
	vocdoniFollowers, err1 := u.fapi.UserFollowers(internalCtx, VocdoniFarcasterFID)
	if err1 == nil {
		for _, fid := range vocdoniFollowers {
			u.vocdoniFollowers[fid] = true
		}
	}
	// update votecaster followers
	votecasterFollowers, err2 := u.fapi.UserFollowers(internalCtx, VotecasterFarcasterFID)
	if err2 == nil {
		for _, fid := range votecasterFollowers {
			u.votecasterFollowers[fid] = true
		}
	}
	// update recasters
	recasters, err3 := u.fapi.RecastsFIDs(internalCtx, &farcasterapi.APIMessage{
		Author: VocdoniFarcasterFID,
		Hash:   VotecasterAnnouncementCastHash,
	})
	if err3 == nil {
		for _, fid := range recasters {
			u.recasters[fid] = true
		}
	}
	if err != nil || err1 != nil || err2 != nil || err3 != nil {
		return fmt.Errorf("error updating internal followers: %w, %w, %w, %w", err, err1, err2, err3)
	}
	return nil
}

// updateHolders method updates the internal holders lists to cache the holders
// of the Votecaster NFT pass, the Votecaster Launch NFT, the KIWI token, the
// DegenDAO NFT, the Haberdashery NFT, the TokyoDAO NFT, the Proxy, the
// ProxyStudio NFT, and the NameDegen NFT. It fetches the holders data from the
// Airstack API and the Census3 API. It returns an error if the holders data
// cannot be fetched.
func (u *Updater) updateHolders() error {
	u.holdersMtx.Lock()
	defer u.holdersMtx.Unlock()
	var errs []error
	// update Votecaster NFT pass holders
	votecasterNFTPassHolders, err := u.airstack.Client.TokenBalances(
		VotecasterNFTPassAddress, VotecasterNFTPassChainShortName)
	if err != nil {
		errs = append(errs, fmt.Errorf("error getting votecaster nft pass holders: %w", err))
	} else {
		u.votecasterLaunchNFTHolders = make(map[common.Address]*big.Int)
		for _, holder := range votecasterNFTPassHolders {
			if holder.Balance.Cmp(big.NewInt(0)) > 0 {
				u.votecasterNFTPassHolders[holder.Address] = holder.Balance
			}
		}
		log.Infow("votecaster nft pass holders", "holders", len(u.votecasterLaunchNFTHolders))
	}
	// update Votecaster Launch NFT holders
	votecasterLaunchNFTHolders, err := u.airstack.Client.TokenBalances(
		VotecasterLaunchNFTAddress, VotecasterLaunchNFTChainShortName)
	if err != nil {
		errs = append(errs, fmt.Errorf("error getting votecaster launch nft holders: %w", err))
	} else {
		u.votecasterLaunchNFTHolders = make(map[common.Address]*big.Int)
		for _, holder := range votecasterLaunchNFTHolders {
			if holder.Balance.Cmp(big.NewInt(0)) > 0 {
				u.votecasterLaunchNFTHolders[holder.Address] = holder.Balance
			}
		}
		log.Infow("votecaster launch nft holders", "holders", len(u.votecasterLaunchNFTHolders))
	}
	// update KIWI holders
	kiwiToken, err := u.census3.Token(KIWIAddress.Hex(), KIWIChainID, "")
	if err != nil {
		errs = append(errs, fmt.Errorf("error getting KIWI token info: %w", err))
	} else {
		kiwiHoldersQueueID, err := u.census3.HoldersByStrategy(kiwiToken.DefaultStrategy)
		if err != nil {
			errs = append(errs, fmt.Errorf("error getting KIWI holders queue ID: %w", err))
		} else {
			u.kiwiHolders = make(map[common.Address]*big.Int)
			for {
				kiwiHolders, finished, err := u.census3.HoldersByStrategyQueue(
					kiwiToken.DefaultStrategy, kiwiHoldersQueueID)
				if err != nil {
					errs = append(errs, fmt.Errorf("error getting KIWI holders: %w", err))
					break
				}
				if finished {
					for holder, balance := range kiwiHolders {
						if balance.Cmp(big.NewInt(0)) > 0 {
							u.kiwiHolders[holder] = balance
						}
					}
					break
				}
				time.Sleep(time.Second * 1)
			}
			log.Infow("KIWI holders", "holders", len(u.kiwiHolders))
		}
	}
	// update DegenDAO NFT holders
	degenDAONFTHolders, err := u.airstack.Client.TokenBalances(
		DegenDAONFTAddress, DegenDAONFTChainShortName)
	if err != nil {
		errs = append(errs, fmt.Errorf("error getting DegenDAO NFT holders: %w", err))
	} else {
		u.degenDAONFTHolders = make(map[common.Address]*big.Int)
		for _, holder := range degenDAONFTHolders {
			if holder.Balance.Cmp(big.NewInt(0)) > 0 {
				u.degenDAONFTHolders[holder.Address] = holder.Balance
			}
		}
		log.Infow("DegenDAO NFT holders", "holders", len(u.degenDAONFTHolders))
	}
	// update Haberdashery NFT holders
	haberdasheryNFTHolders, err := u.airstack.Client.TokenBalances(
		HaberdasheryNFTAddress, HaberdasheryNFTChainShortName)
	if err != nil {
		errs = append(errs, fmt.Errorf("error getting Haberdashery NFT holders: %w", err))
	} else {
		u.haberdasheryNFTHolders = make(map[common.Address]*big.Int)
		for _, holder := range haberdasheryNFTHolders {
			if holder.Balance.Cmp(big.NewInt(0)) > 0 {
				u.haberdasheryNFTHolders[holder.Address] = holder.Balance
			}
		}
		log.Infow("Haberdashery NFT holders", "holders", len(u.haberdasheryNFTHolders))
	}
	// update TokyoDAO NFT holders
	tokyoDAONFTHolders, err := u.airstack.Client.TokenBalances(
		TokyoDAONFTAddress, TokyoDAONFTChainShortName)
	if err != nil {
		errs = append(errs, fmt.Errorf("error getting TokyoDAO NFT holders: %w", err))
	} else {
		u.tokyoDAONFTHolders = make(map[common.Address]*big.Int)
		for _, holder := range tokyoDAONFTHolders {
			if holder.Balance.Cmp(big.NewInt(0)) > 0 {
				u.tokyoDAONFTHolders[holder.Address] = holder.Balance
			}
		}
		log.Infow("TokyoDAO NFT holders", "holders", len(u.tokyoDAONFTHolders))
	}
	// update Proxy
	proxyHolders, err := u.airstack.Client.TokenBalances(
		ProxyAddress, ProxyChainShortName)
	if err != nil {
		errs = append(errs, fmt.Errorf("error getting Proxy holders: %w", err))
	} else {
		u.proxyHolders = make(map[common.Address]*big.Int)
		for _, holder := range proxyHolders {
			if holder.Balance.Cmp(big.NewInt(0)) > 0 {
				u.proxyHolders[holder.Address] = holder.Balance
			}
		}
		log.Infow("Proxy holders", "holders", len(u.proxyHolders))
	}
	// update ProxyStudio NFT holders
	proxyStudioNFTHolders, err := u.airstack.Client.TokenBalances(
		ProxyStudioNFTAddress, ProxyStudioNFTShortName)
	if err != nil {
		errs = append(errs, fmt.Errorf("error getting ProxyStudio NFT holders: %w", err))
	} else {
		u.proxyStudioNFTHolders = make(map[common.Address]*big.Int)
		for _, holder := range proxyStudioNFTHolders {
			if holder.Balance.Cmp(big.NewInt(0)) > 0 {
				u.proxyStudioNFTHolders[holder.Address] = holder.Balance
			}
		}
		log.Infow("ProxyStudio NFT holders", "holders", len(u.proxyStudioNFTHolders))
	}
	// update NameDegen NFT holders
	nameDegenHolders, err := u.airstack.Client.TokenBalances(
		NameDegenAddress, NameDegenChainShortName)
	if err != nil {
		errs = append(errs, fmt.Errorf("error getting NameDegen NFT holders: %w", err))
	} else {
		u.nameDegenHolders = make(map[common.Address]*big.Int)
		for _, holder := range nameDegenHolders {
			if holder.Balance.Cmp(big.NewInt(0)) > 0 {
				u.nameDegenHolders[holder.Address] = holder.Balance
			}
		}
		log.Infow("NameDegen NFT holders", "holders", len(u.nameDegenHolders))
	}
	if len(errs) > 0 {
		return fmt.Errorf("error updating holders: %v", errs)
	}
	return nil
}

// updateUsers method iterates over all users in the database and updates their
// reputation data. It uses a concurrent approach to update multiple users at
// the same time, limiting the number of concurrent updates to the maximum
// number of concurrent updates set in the Updater instance. It fetches the
// activity data from the database and the boosters data from the Airstack and
// the Census3 API.
func (u *Updater) updateUsers() error {
	log.Info("updating users reputation")
	ctx, cancel := context.WithCancel(u.ctx)
	defer cancel()
	// limit the number of concurrent updates and create the channel to receive
	// the users, creates also the inner waiter to wait for all updates to
	// finish
	concurrentUpdates := make(chan struct{}, u.maxConcurrent)
	usersChan := make(chan *mongo.User)
	innerWaiter := sync.WaitGroup{}
	// counters for total and updated users
	updates := atomic.Int64{}
	total := atomic.Int64{}
	// listen for users and update them concurrently
	innerWaiter.Add(1)
	go func() {
		defer innerWaiter.Done()
		for user := range usersChan {
			select {
			case <-ctx.Done():
				return
			default:
				total.Add(1)
				// get a slot in the concurrent updates channel
				concurrentUpdates <- struct{}{}
				go func(user *mongo.User) {
					// release the slot when the update is done
					defer func() {
						<-concurrentUpdates
					}()
					// update user reputation
					if err := u.updateUser(user); err != nil {
						log.Errorf("error updating user %d: %v", user.UserID, err)
					} else {
						updates.Add(1)
					}
				}(user)
			}
		}
	}()
	// iterate over users and send them to the channel
	if err := u.db.UsersIterator(ctx, usersChan); err != nil {
		return fmt.Errorf("error iterating users: %w", err)
	}
	innerWaiter.Wait()
	log.Infow("users reputation updated", "total", total.Load(), "updated", updates.Load())
	return nil
}

func (u *Updater) updateCommunities() error {
	log.Info("updating communities reputation")
	// limit the number of concurrent updates and create the channel to receive
	// the communities, creates also the inner waiter to wait for all updates to
	// finish
	concurrentUpdates := make(chan struct{}, u.maxConcurrent)
	innerWaiter := sync.WaitGroup{}
	communities, total, err := u.db.ListCommunities(-1, 0)
	if err != nil {
		return fmt.Errorf("error listing communities: %w", err)
	}
	// counters for total and updated communities
	updates := atomic.Int64{}
	// listen for communities and update them concurrently
	innerWaiter.Add(1)
	go func() {
		defer innerWaiter.Done()
		for _, community := range communities {
			// get a slot in the concurrent updates channel
			concurrentUpdates <- struct{}{}
			go func(community *mongo.Community) {
				// release the slot when the update is done
				defer func() {
					<-concurrentUpdates
				}()
				participation, censusSize, err := u.communityPoints(community)
				if err != nil {
					log.Errorf("error getting community %s points: %v", community.ID, err)
					return
				}
				if err := u.db.SetCommunityPoints(community.ID, participation, censusSize); err != nil {
					log.Errorf("error updating community %s reputation: %v", community.ID, err)
					return
				}
				updates.Add(1)
			}(&community)
		}
	}()
	innerWaiter.Wait()
	log.Infow("communities reputation updated", "total", total, "updated", updates.Load())
	return nil
}

// updateUser method updates the reputation data of a given user. It fetches the
// activity data from the database and the boosters data from the Airstack and
// the Census3 API. It then updates the reputation data in the database.
func (u *Updater) updateUser(user *mongo.User) error {
	if u.db == nil {
		return fmt.Errorf("database not set")
	}
	rep, err := u.db.DetailedUserReputation(user.UserID)
	if err != nil {
		// if the user is not found, create a new user with blank data
		if errors.Is(err, mongo.ErrUserUnknown) {
			return u.db.SetDetailedReputationForUser(user.UserID, &mongo.Reputation{})
		}
		// return the error if it is not a user unknown error
		return err
	}
	// get activiy data if needed
	activityRep, err := u.userActivityReputation(user)
	if err != nil {
		// if there is an error fetching the activity data, log the error and
		// continue updating the no failed activity data
		log.Warnw("error getting user activity reputation", "error", err, "user", user.UserID)
	}
	// update reputation
	rep.FollowersCount = activityRep.FollowersCount
	rep.ElectionsCreatedCount = activityRep.ElectionsCreatedCount
	rep.CastVotesCount = activityRep.CastVotesCount
	rep.ParticipationsCount = activityRep.ParticipationsCount
	rep.CommunitiesCount = activityRep.CommunitiesCount
	// get boosters data if needed
	boostersRep := u.userBoosters(user)
	// if there is an error fetching the boosters data, log the error and
	// continue updating the no failed boosters data
	if err != nil {
		log.Warnw("error getting some boosters", "error", err, "user", user.UserID)
	}
	// update reputation
	rep.HasVotecasterNFTPass = boostersRep.HasVotecasterNFTPass
	rep.HasVotecasterLaunchNFT = boostersRep.HasVotecasterLaunchNFT
	rep.IsVotecasterAlphafrensFollower = boostersRep.IsVotecasterAlphafrensFollower
	rep.IsVotecasterFarcasterFollower = boostersRep.IsVotecasterFarcasterFollower
	rep.IsVocdoniFarcasterFollower = boostersRep.IsVocdoniFarcasterFollower
	rep.VotecasterAnnouncementRecasted = boostersRep.VotecasterAnnouncementRecasted
	rep.HasKIWI = boostersRep.HasKIWI
	rep.HasDegenDAONFT = boostersRep.HasDegenDAONFT
	rep.Has10kDegenAtLeast = boostersRep.Has10kDegenAtLeast
	rep.HasTokyoDAONFT = boostersRep.HasTokyoDAONFT
	rep.Has5ProxyAtLeast = boostersRep.Has5ProxyAtLeast
	rep.HasNameDegen = boostersRep.HasNameDegen
	// commit reputation
	return u.db.SetDetailedReputationForUser(user.UserID, rep)
}

// userActivityReputation method fetches the activity data of a given user from
// the database. It returns the activity data as an ActivityReputation struct.
// The activity data includes the number of followers, the number of elections
// created, the number of casted votes, the number of votes casted on elections
// created by the user, and the number of communities where the user is an
// admin. It returns an error if the activity data cannot be fetched.
func (u *Updater) userActivityReputation(user *mongo.User) (*ActivityReputationCounts, error) {
	// Fetch the total votes cast on elections created by the user
	totalVotes, err := u.db.TotalVotesForUserElections(user.UserID)
	if err != nil {
		return &ActivityReputationCounts{}, fmt.Errorf("error fetching total votes for user elections: %w", err)
	}
	// Fetch the number of communities where the user is an admin
	communitiesCount, err := u.db.CommunitiesCountForUser(user.UserID)
	if err != nil {
		return &ActivityReputationCounts{}, fmt.Errorf("error fetching communities count for user: %w", err)
	}
	return &ActivityReputationCounts{
		FollowersCount:        user.Followers,
		ElectionsCreatedCount: user.ElectionCount,
		CastVotesCount:        user.CastedVotes,
		ParticipationsCount:   totalVotes,
		CommunitiesCount:      communitiesCount,
	}, nil
}

func (u *Updater) communityPoints(community *mongo.Community) (float64, uint64, error) {
	participation, err := u.db.CommunityParticipationMean(community.ID)
	if err != nil {
		return 0, 0, fmt.Errorf("error fetching community participation mean: %w", err)
	}
	ctx, cancel := context.WithTimeout(u.ctx, time.Second*30)
	defer cancel()
	var censusSize uint64
	switch community.Census.Type {
	case mongo.TypeCommunityCensusChannel:
		users, err := u.fapi.ChannelFIDs(ctx, community.Census.Channel, nil)
		if err != nil {
			return 0, 0, fmt.Errorf("error fetching channel users: %w", err)
		}
		censusSize = uint64(len(users))
	case mongo.TypeCommunityCensusERC20, mongo.TypeCommunityCensusNFT:
		singleUsers := map[common.Address]bool{}
		for _, token := range community.Census.Addresses {
			holders, err := u.airstack.TokenBalances(common.HexToAddress(token.Address), token.Blockchain)
			if err != nil {
				return 0, 0, fmt.Errorf("error fetching token holders: %w", err)
			}
			for _, holder := range holders {
				if _, ok := singleUsers[holder.Address]; !ok {
					singleUsers[holder.Address] = true
				}
			}
		}
		censusSize = uint64(len(singleUsers))
	case mongo.TypeCommunityCensusFollowers:
		fid, err := communityhub.DecodeUserChannelFID(community.Census.Channel)
		if err != nil {
			return 0, 0, fmt.Errorf("invalid follower census user reference: %w", err)
		}
		users, err := u.fapi.UserFollowers(ctx, fid)
		if err != nil {
			return 0, 0, fmt.Errorf("error fetching user followers: %w", err)
		}
		censusSize = uint64(len(users))
	default:
		return 0, 0, fmt.Errorf("invalid census type")
	}
	return participation, censusSize, nil
}

// userBoosters method fetches the boosters data of a given user from the
// Airstack and the Census3 API. It returns the boosters data as a Boosters
// struct. The boosters data includes whether the user has the Votecaster NFT
// pass, the Votecaster Launch NFT, the user is subscribed to Votecaster
// Alphafrens channel, the user follows Votecaster and Vocdoni profiles on
// Farcaster, the user has recasted the Votecaster Launch cast announcement,
// the user has KIWI, the user has the DegenDAO NFT, the user has at least 10k
// Degen, the user has Haberdashery NFT, the user has the TokyoDAO NFT, the user
// has a Proxy, the user has at least 5 Proxies, the user has the ProxyStudio
// NFT, and the user has the NameDegen NFT. It returns an error if the boosters
// data cannot be fetched.
func (u *Updater) userBoosters(user *mongo.User) *Boosters {
	// create new boosters struct and slice for errors
	boosters := &Boosters{}
	// check if user is votecaster alphafrens follower, is vocdoni or votecaster
	// farcaster follower, and if the user has recasted the votecaster launch
	// cast announcement
	u.followersMtx.Lock()
	defer u.followersMtx.Unlock()
	boosters.IsVotecasterAlphafrensFollower = u.alfafrensFollowers[user.UserID]
	boosters.IsVocdoniFarcasterFollower = u.vocdoniFollowers[user.UserID]
	boosters.IsVotecasterFarcasterFollower = u.votecasterFollowers[user.UserID]
	boosters.VotecasterAnnouncementRecasted = u.recasters[user.UserID]
	// for every user address check every booster only if it is not already set
	u.holdersMtx.Lock()
	defer u.holdersMtx.Unlock()
	log.Info(user.Addresses)
	for _, strAddr := range user.Addresses {
		addr := common.HexToAddress(strAddr)
		// check if user has votecaster nft pass
		if !boosters.HasVotecasterNFTPass {
			_, ok := u.votecasterNFTPassHolders[addr]
			boosters.HasVotecasterNFTPass = ok
		}
		// check if user has votecaster launch nft
		if !boosters.HasVotecasterLaunchNFT {
			_, ok := u.votecasterLaunchNFTHolders[addr]
			boosters.HasVotecasterLaunchNFT = ok
		}
		// check if user has KIWI
		if !boosters.HasKIWI {
			_, ok := u.kiwiHolders[addr]
			boosters.HasKIWI = ok
		}
		// check if user has DegenDAO NFT
		if !boosters.HasDegenDAONFT {
			_, ok := u.degenDAONFTHolders[addr]
			boosters.HasDegenDAONFT = ok
		}
		// check if user has Haberdashery NFT
		if !boosters.HasHaberdasheryNFT {
			_, ok := u.haberdasheryNFTHolders[addr]
			boosters.HasHaberdasheryNFT = ok
		}
		// check if user has 10k Degen
		if !boosters.Has10kDegenAtLeast {
			if balance, ok := u.degenDAONFTHolders[addr]; ok {
				boosters.Has10kDegenAtLeast = balance.Cmp(big.NewInt(10000)) >= 0
			}
		}
		// check if user has TokyoDAO NFT
		if !boosters.HasTokyoDAONFT {
			_, ok := u.tokyoDAONFTHolders[addr]
			boosters.HasTokyoDAONFT = ok
		}
		// check if user has Proxy and at least 5 Proxies
		if !boosters.Has5ProxyAtLeast {
			if balance, ok := u.proxyHolders[addr]; ok {
				boosters.Has5ProxyAtLeast = balance.Cmp(big.NewInt(5)) >= 0
			}
		}
		// check if user has ProxyStudio NFT
		if !boosters.HasProxyStudioNFT {
			_, ok := u.proxyStudioNFTHolders[addr]
			boosters.HasProxyStudioNFT = ok
		}
		// check if user has NameDegen
		if !boosters.HasNameDegen {
			_, ok := u.nameDegenHolders[addr]
			boosters.HasNameDegen = ok
		}
	}
	return boosters
}
