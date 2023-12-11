package outbox

import (
	"fmt"

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

func makeOutboundRequest(outbound OutboundRequest) {
	// send request back to slack

	fmt.Printf(`outbound request: [caller:%s] [response:%s]`, outbound.Caller, outbound.Response)

	// api := slack.New("!!TOKEN!!")
	api := slack.New("")

	/*groups, err := api.GetUserGroups(slack.GetUserGroupsOptionIncludeUsers(false))

	if err != nil {
		fmt.Printf("%s\n", err)
		return
	}

	for _, group := range groups {
		fmt.Printf("ID: %s, Name: %s\n", group.ID, group.Name)
	}*/

	// user, err := api.GetUserByEmail("")
	/*user, err := api.GetUserInfo("")
	if err != nil {
		fmt.Printf("%s\n", err)
		return
	}
	fmt.Printf("ID: %s, Fullname: %s, Email: %s\n", user.ID, user.Profile.RealName, user.Profile.Email)*/

	api.PostMessage("", slack.MsgOptionText(outbound.Response, false))
}
