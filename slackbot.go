package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"slackbot/inbox"
	"slackbot/outbox"
	router "slackbot/router"
)

func main() {

	port, err := strconv.Atoi(os.Getenv("SERVICE_PORT"))
	if err != nil {
		port = 5001
	}

	servePort := fmt.Sprintf(":%d", port)

	rtr := router.GetRouter()
	inbox.ProcessInbox()
	outbox.ProcessOubox()

	fmt.Printf("Listening on port %s", servePort)
	log.Fatal(http.ListenAndServe(servePort, rtr))
}
