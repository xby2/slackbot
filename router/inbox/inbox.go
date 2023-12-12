package inbox

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"slackbot/inbox"
	"slackbot/outbox"
)

func InboxRouteHandler(w http.ResponseWriter, r *http.Request) {
	var incomingRequest inbox.InboundRequest

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&incomingRequest); err != nil {
		resp, _ := json.Marshal("invalid request")
		w.WriteHeader(500)
		w.Write(resp)
		return
	}

	// create request and send to subroutine for processing
	inbound := inbox.InboundRequest{
		Caller:  incomingRequest.Caller,
		Message: incomingRequest.Message,
		Start:   time.Now(),
	}

	inbox.Inbox <- inbound

	// send in progress message to use
	outbound := outbox.OutboundRequest{
		Caller:   incomingRequest.Caller,
		Message:  incomingRequest.Message,
		Response: "Request processing",
		Start:    time.Now(),
		End:      time.Now(),
	}

	outbox.Outbox <- outbound

	w.WriteHeader(200)
	w.Write([]byte(fmt.Sprintf(`{ "status": "received", "caller": "%s", "message": "%s" }`, inbound.Caller, inbound.Message)))
}
