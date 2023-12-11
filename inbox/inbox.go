package inbox

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"time"

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
	// var j = []byte(fmt.Sprintf(`{"prompt": "You are a helpful assistant.", "message": "%s"}`, inbound.Message))
	var j = []byte(fmt.Sprintf(`{"prompt": "You are a proficient AI with a specialty in distilling information into key points. Based on the following text, identify and list the main points that were discussed or brought up. These should be the most important ideas, findings, or topics that are crucial to the essence of the discussion. Your goal is to provide a list that someone could read to quickly understand what was talked about.", "message": "%s"}`, inbound.Message))

	req, _ := http.NewRequest("POST", "http://localhost:5005/eksy", bytes.NewBuffer(j))
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{
		Timeout: 5 * time.Minute,
	}
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
