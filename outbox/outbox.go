package outbox

import (
	"fmt"
	"os"

	"github.com/slack-go/slack"
)

var Outbox = make(chan OutboundRequest, 100)

func ProcessOubox() {
	go func() {
		for {
			select {
			case outbound := <-Outbox:
				makeOutboundRequest(outbound)
			}
		}
	}()
}

var SLACKBOT_TOKEN string = os.Getenv("SLACKBOT_TOKEN")

func makeOutboundRequest(outbound OutboundRequest) {
	// send request back to slack

	api := slack.New(SLACKBOT_TOKEN)

	// user, err := api.GetUserByEmail("")

	msg := fmt.Sprintf("Elapsed Time[%s]\nMessage[%s]\n\nResponse\n%s", outbound.End.Sub(outbound.Start), outbound.Message, outbound.Response)
	api.PostMessage(outbound.Caller, slack.MsgOptionText(msg, false))
}
