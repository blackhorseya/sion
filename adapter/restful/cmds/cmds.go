package cmds

import (
	"github.com/blackhorseya/sion/pkg/contextx"
	"github.com/line/line-bot-sdk-go/v8/linebot/messaging_api"
)

// TextCommander is the interface for text command.
type TextCommander interface {
	Execute(ctx contextx.Contextx, text string) ([]messaging_api.MessageInterface, error)
}

// NewCommands is a function to create a new commands.
func NewCommands() []TextCommander {
	return []TextCommander{
		&PingCommand{},
	}
}
