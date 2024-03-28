package cmds

import (
	"github.com/blackhorseya/sion/pkg/contextx"
	"github.com/line/line-bot-sdk-go/v8/linebot/messaging_api"
)

// PingCommand is a struct to handle ping command.
type PingCommand struct {
}

func (cmd *PingCommand) Execute(ctx contextx.Contextx, text string) ([]messaging_api.MessageInterface, error) {
	if text == "ping" {
		return []messaging_api.MessageInterface{
			&messaging_api.TextMessage{
				Text: "pong",
			},
		}, nil
	}

	return nil, nil
}
