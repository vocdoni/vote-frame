package bot

import "fmt"

var (
	// ErrAPINotSet is returned when the API is not set in the bot configuration.
	ErrAPINotSet = fmt.Errorf("api not set")
	// ErrParsingPoll is returned when there is an error parsing the poll during
	// the poll message handler.
	ErrParsingPoll = fmt.Errorf("error parsing poll")
	// ErrGettingUserData is returned when there is an error getting user data
	// during the poll message handler.
	ErrGettingUserData = fmt.Errorf("error getting user data")
	// ErrReplyingToCast is returned when there is an error replying to a cast
	// during the reply with poll URL function.
	ErrReplyingToCast = fmt.Errorf("error replying to cast")
	// ErrGettingParentCast is returned when there is an error getting the parent
	// cast during the mute notifications message handler.
	ErrGettingParentCast = fmt.Errorf("error getting parent cast")
)
