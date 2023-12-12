package inbox

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"

	"slackbot/outbox"
	"slackbot/prompt"
)

var Inbox = make(chan InboundRequest, 100)

var SERVICE_ADDRESS string = os.Getenv("SERVICE_ADDRESS")
var BOT_TYPE string = os.Getenv("BOT_TYPE")

func ProcessInbox() {
	go func() {
		for {
			select {
			case inbound := <-Inbox:
				makeInboundRequest(inbound)
			}
		}
	}()
}

func makeInboundRequest(inbound InboundRequest) {
	// make request to chatbot

	// prepare Body for request, includes: prompt, message, maxtokens
	// prompt pulled from prompt package
	var j = []byte(fmt.Sprintf(`{"prompt": "%s", "message": "%s", "maxtokens": 100}`, prompt.PROMPTS[BOT_TYPE], inbound.Message))

	// make request, timeout set to 5 minutes
	req, _ := http.NewRequest("POST", SERVICE_ADDRESS, bytes.NewBuffer(j))
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{
		Timeout: 5 * time.Minute,
	}
	resp, err := client.Do(req)

	// on error, push error response to outbox
	if err != nil {
		outbound := outbox.OutboundRequest{
			Caller:   inbound.Caller,
			Message:  inbound.Message,
			Response: err.Error(),
			Start:    inbound.Start,
			End:      time.Now(),
		}

		outbox.Outbox <- outbound
		return
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)

	outbound := outbox.OutboundRequest{
		Caller:   inbound.Caller,
		Message:  inbound.Message,
		Response: string(body),
		Start:    inbound.Start,
		End:      time.Now(),
	}

	// push response to outbox
	outbox.Outbox <- outbound
	//fmt.Printf(`inbound request: [caller:%s] [message:%s]`, inbound.Caller, inbound.Message)
}
