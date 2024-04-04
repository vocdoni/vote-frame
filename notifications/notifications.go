package notifications

import (
	"context"
	"errors"
	"fmt"
	"sync"
	"time"

	"github.com/vocdoni/vote-frame/farcasterapi"
	"github.com/vocdoni/vote-frame/mongo"
	"go.vocdoni.io/dvote/log"
)

const (
	DefaultListenCoolDown       = 30 * time.Second
	DefaultSendCoolDown         = 500 * time.Millisecond
	DefaultNotificationDeadline = 24 * time.Hour
	DefaultPermissionMessage    = `👋 Hey @%s! 

I'm the farcaster.vote bot. You're included in a poll census created by %s, but I won't bother you again if you prefer not to receive notifications.

Please let me know if you want to be notified or not! 🤖👍`
	DefaultNotificationMessage = `👋 Hey @%s!

The user %s created a new poll!

🗳 And you're eligible to vote!

Cast your vote to make a difference 👇.

(to mute this user reply to this message with: @%s mute)`
)

// notificationThread is the parent cast to reply to when sending a notification
// and avoid spamming the account feed. https://warpcast.com/vocdoni/0xfd847188
var notificationThread = &farcasterapi.APIMessage{
	Hash:   "0xfd8471884f3aaf3528d33ba8ae59f57904124d27",
	Author: 7548,
}

type NotifificationManagerConfig struct {
	DB                   *mongo.MongoStorage
	API                  farcasterapi.API
	ListenCoolDown       time.Duration
	DefaultSendCoolDown  time.Duration
	NotificationDeadline time.Duration
	PermissionMessage    string
	NotificationMessage  string
	FrameURL             string
}

// NotificationManager is a manager that listens for new notifications registered
// in the database and sends them to the users via the farcaster API.
type NotificationManager struct {
	ctx             context.Context
	cancel          context.CancelFunc
	db              *mongo.MongoStorage
	api             farcasterapi.API
	listenCoolDown  time.Duration
	sendCoolDown    time.Duration
	permissionMsg   string
	notificationMsg string
	frameURL        string
}

// check method checks the configuration required values and sets default values
// for optional configuration values if not provided. It returns nil if the
// configuration is correct and error if not.
func (conf *NotifificationManagerConfig) check() error {
	// check required configuration values
	if conf.DB == nil {
		return fmt.Errorf("database is required")
	}
	if conf.API == nil {
		return fmt.Errorf("farcaster API is required")
	}
	if conf.FrameURL == "" {
		return fmt.Errorf("frame URL is required")
	}
	// check optional configuration values and set default values if not provided
	if conf.ListenCoolDown == 0 {
		conf.ListenCoolDown = DefaultListenCoolDown
	}
	if conf.DefaultSendCoolDown == 0 {
		conf.DefaultSendCoolDown = DefaultSendCoolDown
	}
	if conf.NotificationDeadline == 0 {
		conf.NotificationDeadline = DefaultNotificationDeadline
	}
	if conf.PermissionMessage == "" {
		conf.PermissionMessage = DefaultPermissionMessage
	}
	if conf.NotificationMessage == "" {
		conf.NotificationMessage = DefaultNotificationMessage
	}
	return nil
}

// New creates a new NotificationManager instance with the given context, database
// and farcaster API. It also sets the listen cool down duration.
func New(ctx context.Context, config *NotifificationManagerConfig) (*NotificationManager, error) {
	if err := config.check(); err != nil {
		return nil, err
	}
	ctx, cancel := context.WithCancel(ctx)
	return &NotificationManager{
		ctx:             ctx,
		cancel:          cancel,
		db:              config.DB,
		api:             config.API,
		listenCoolDown:  config.ListenCoolDown,
		sendCoolDown:    config.DefaultSendCoolDown,
		permissionMsg:   config.PermissionMessage,
		notificationMsg: config.NotificationMessage,
		frameURL:        config.FrameURL,
	}, nil
}

// Start starts the notification manager and listens for new notifications in the
// database to send them to the users. It uses a cool down duration to avoid
// spamming the farcaster API. It runs in the background and send notifications
// in parallel.
func (nm *NotificationManager) Start() {
	go func() {
		for {
			select {
			case <-nm.ctx.Done():
				return
			case <-time.After(nm.listenCoolDown):
				notifications, err := nm.db.LastNotifications(100)
				if err != nil {
					log.Errorf("error getting notifications: %s", err)
					continue
				}
				log.Infow("notifications found", "count", len(notifications))
				if err := nm.handleNotifications(notifications); err != nil {
					log.Errorf("error sending notifications: %s", err)
				}
			}
		}
	}()
}

// Stop stops the notification manager and cancels the context.
func (nm *NotificationManager) Stop() {
	nm.cancel()
}

// handleNotifications sends the notifications to the users and removes them
// from the database. It uses a semaphore to limit the number of concurrent
// goroutines and an error channel to return any error found. It checks if the
// user to notify has accepted the notifications, if not, requests the
// permission. It also purges the notifications that have not been accepted
// after its deadline.
func (nm *NotificationManager) handleNotifications(notifications []mongo.Notification) error {
	// create channels and waitgroup, the semaphore is used to limit the number
	// of concurrent goroutines and the error channel is used to return any
	// error found
	sem := make(chan struct{}, 10)
	errCh := make(chan error, 1)
	wg := sync.WaitGroup{}
	// iterate over notifications and send them
	for _, n := range notifications {
		// add goroutine to waitgroup and semaphore
		wg.Add(1)
		sem <- struct{}{}
		go func(n mongo.Notification) {
			defer wg.Done()
			defer func() { <-sem }()
			allowed, err := nm.checkOrReqPermission(n.UserID, n.Username, n.AuthorUsername)
			if err != nil {
				if errors.Is(err, mongo.ErrUserUnknown) {
					log.Debugw("user not found", "user", n.UserID)
					if err := nm.db.RemoveNotification(n.ID); err != nil {
						errCh <- fmt.Errorf("error deleting notification: %s", err)
					}
					return
				}
				errCh <- fmt.Errorf("error checking or requesting permission: %s", err)
				return
			}
			// if the user has not accepted the notifications, check if the deadline
			// has been reached and remove the notification if so, or continue to the
			// next notification
			if !allowed {
				// if the user has not accepted the notifications yet, the
				// notification have non zero deadline, if that deadline is
				// reached, and the notification is not accepted, the
				// notification must be removed from the database. If the
				// deadline is zero, the notification permission has been
				// denied, and the notification must be removed
				if time.Now().After(n.Deadline) {
					log.Debugw("notification deadline reached, purging...", "notification", n.ID)
					if err := nm.db.RemoveNotification(n.ID); err != nil {
						errCh <- fmt.Errorf("error deleting notification: %s", err)
					}
				}
				return
			}
			// check if the receiver has muted the creator of the notification
			isCreatorMuted, err := nm.db.IsUserNotificationMuted(n.UserID, n.AuthorID)
			if err != nil {
				errCh <- fmt.Errorf("error checking if user is muted: %s", err)
				return
			}
			// if the creator is muted by the receiver, remove the notification
			// from the database
			if isCreatorMuted {
				log.Debugw("creator muted by notification receiver, purging notification...",
					"notification", n.ID,
					"receiver", n.UserID,
					"creator", n.AuthorID)
				if err := nm.db.RemoveNotification(n.ID); err != nil {
					errCh <- fmt.Errorf("error deleting notification: %s", err)
				}
				return
			}
			// retrieve the current user data from the API
			ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
			defer cancel()
			userdata, err := nm.api.UserDataByFID(ctx, nm.api.FID())
			if err != nil {
				errCh <- fmt.Errorf("error retrieving bot user data: %s", err)
				return
			}
			log.Infof("%+v", userdata)
			// send notification and remove it from the database
			log.Debugw("permission granted, sending and removing notification...", "notification", n.ID)
			msg := fmt.Sprintf(nm.notificationMsg, n.Username, n.AuthorUsername, userdata.Username)
			if err := nm.api.Reply(nm.ctx, notificationThread, msg, []uint64{n.UserID}, n.FrameUrl); err != nil {
				errCh <- fmt.Errorf("error sending notification: %s", err)
				return
			}
			if err := nm.db.RemoveNotification(n.ID); err != nil {
				errCh <- fmt.Errorf("error deleting notification: %s", err)
				return
			}
		}(n)
	}
	// wait for all goroutines to finish and close channels
	go func() {
		wg.Wait()
		close(errCh)
		close(sem)
	}()
	// listen error channel and return any err error found
	for err := range errCh {
		return err
	}
	return nil
}

// checkOrReqPermission checks if the user has accepted the notifications, if not,
// it sends a notification request with the permission message and the frame URL.
// It also updates the access profile with the notification requested status. If
// the user has not accepted the notifications, it returns false, otherwise, it
// returns true. If an error occurs, it returns the error.
func (nm *NotificationManager) checkOrReqPermission(userID uint64, username, authorUsername string) (bool, error) {
	alreadyRequested := false

	profile, err := nm.db.UserAccessProfile(userID)
	if err != nil {
		if !errors.Is(err, mongo.ErrUserUnknown) {
			return false, err
		}
	} else {
		alreadyRequested = profile.NotificationsRequested
	}
	// if the user has requested notifications, return the accepted status
	if alreadyRequested {
		log.Debugw("notifications requested",
			"user", userID,
			"granted", profile.NotificationsAccepted)
		return profile.NotificationsAccepted, nil
	}
	log.Debugw("notifications not requested, requesting...", "user", userID)
	// if the user has not been requested for notifications yet, send the
	// notification request with the permission message and the frame URL
	msg := fmt.Sprintf(nm.permissionMsg, username, authorUsername)
	if err := nm.api.Publish(nm.ctx, msg, []uint64{userID}, nm.frameURL); err != nil {
		return false, fmt.Errorf("error sending notification request: %s", err)
	}
	// update the access profile with the notification requested status
	if err := nm.db.SetNotificationsRequestedForUser(userID, true); err != nil {
		return false, fmt.Errorf("error setting user notification requested: %s", err)
	}
	return false, nil
}
