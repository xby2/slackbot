package main

import (
	"fmt"
	"log"
	"net/http"

	"slackbot/inbox"
	"slackbot/outbox"
	router "slackbot/router"
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
}
