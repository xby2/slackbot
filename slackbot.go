package main

import (
	"fmt"
	"log"
	"net/http"

	"slackbot/inbox"
	"slackbot/outbox"
	router "slackbot/router"
	//"github.com/slack-go/slack"
)

var port int

func main() {

	port = 5001
	servePort := fmt.Sprintf(":%d", port)

	rtr := router.GetRouter()
	inbox.ProcessInbox()
	outbox.ProcessOubox()

	fmt.Printf("Listening on port %s", servePort)
	log.Fatal(http.ListenAndServe(servePort, rtr))

	//api := slack.New("!!TOKEN!!")
	// api := slack.New("!!TOKEN!!")

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
	fmt.Printf("ID: %s, Fullname: %s, Email: %s\n", user.ID, user.Profile.RealName, user.Profile.Email)

	api.PostMessage("UC72G0ATD", slack.MsgOptionText("lookout", false))*/
}
