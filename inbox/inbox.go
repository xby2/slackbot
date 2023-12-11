package inbox

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"slackbot/outbox"
)

var Inbox = make(chan InboundRequest, 100)

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

	// j, _ := json.Marshal(inbound)
	var j = []byte(fmt.Sprintf(`{"prompt": "You are a helpful assistant.", "message": "%s"}`, inbound.Message))

	req, _ := http.NewRequest("POST", "http://localhost:5005/eksy", bytes.NewBuffer(j))
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		outbound := outbox.OutboundRequest{
			Caller:   inbound.Caller,
			Message:  inbound.Message,
			Response: "ERROR_OCCURED",
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
	}

	outbox.Outbox <- outbound

	fmt.Printf(`inbound request: [caller:%s] [message:%s]`, inbound.Caller, inbound.Message)
}
