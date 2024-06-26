package bot

import (
	"context"
	_ "embed"
	"errors"
	"fmt"
	"time"

	"github.com/vocdoni/vote-frame/farcasterapi"
	"go.vocdoni.io/dvote/log"
)

// defaultCoolDown is the default time to wait between casts
const defaultCoolDown = time.Second * 10

// muteRequestContent is the content of a cast that requests to mute a user that
// created a poll and wants to avoid notifications from that user.
const muteRequestContent = "mute"

// BotConfig is the configuration definition for the bot, it includes the API
// instance and the cool down time between casts (default is 10 seconds)
type BotConfig struct {
	API      farcasterapi.API
	CoolDown time.Duration
}

// Bot struct represents a bot that listens for new casts and sends them to a
// channel, it also has a cool down time to avoid spamming the API and a last
// cast timestamp to retrieve new casts from that point, ensuring no cast is
// missed or duplicated
type Bot struct {
	UserData *farcasterapi.Userdata
	api      farcasterapi.API
	ctx      context.Context
	cancel   context.CancelFunc
	coolDown time.Duration
	lastCast uint64
	Messages chan *farcasterapi.APIMessage
}

// New function creates a new bot with the given configuration, it returns an
// error if the API is not set in the configuration.
func New(config BotConfig) (*Bot, error) {
	if config.API == nil {
		return nil, ErrAPINotSet
	}
	if config.CoolDown == 0 {
		config.CoolDown = defaultCoolDown
	}
	bot := &Bot{
		api:      config.API,
		coolDown: config.CoolDown,
		lastCast: uint64(time.Now().Unix()),
		Messages: make(chan *farcasterapi.APIMessage),
	}
	// retrieve the bot user data from the API
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	var err error
	if bot.UserData, err = bot.api.UserDataByFID(ctx, bot.api.FID()); err != nil {
		return nil, fmt.Errorf("error retrieving bot user data: %w", err)
	}
	return bot, nil
}

// Start function starts the bot, it listens for new casts and sends them to the
// Messages channel. It does this in a goroutine to avoid blocking the main and
// every cool down time.
func (b *Bot) Start(ctx context.Context) {
	b.ctx, b.cancel = context.WithCancel(ctx)
	go func() {
		ticker := time.NewTicker(b.coolDown)
		defer ticker.Stop()
		for {
			select {
			case <-b.ctx.Done():
				return
			default:
				// retrieve new messages from the last cast
				messages, lastCast, err := b.api.LastMentions(b.ctx, b.lastCast)
				if err != nil && !errors.Is(err, farcasterapi.ErrNoNewCasts) {
					log.Errorw(err, "error retrieving new casts")
					continue
				}
				b.lastCast = lastCast
				if len(messages) > 0 {
					for _, msg := range messages {
						b.Messages <- msg
					}
				}
				// wait for the cool down time
				<-ticker.C
			}
		}
	}()
}

// Stop function stops the bot and its goroutine, and closes the Messages channel.
func (b *Bot) Stop() {
	if err := b.api.Stop(); err != nil {
		log.Errorf("error stopping bot: %s", err)
	}
	b.cancel()
	close(b.Messages)
}
