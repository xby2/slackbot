package inbox

import (
	"encoding/json"
	"fmt"
	"net/http"

	"slackbot/inbox"
)

const (
	SUCCESS_MESSAGE = "SUCCESS"
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

	inbound := inbox.InboundRequest{
		Caller:  incomingRequest.Caller,
		Message: incomingRequest.Message,
	}

	inbox.Inbox <- inbound

	w.WriteHeader(200)
	// w.Write([]byte(`{ "message": "%s" }`))
	w.Write([]byte(fmt.Sprintf(`{ "status": "received", "caller": "%s", "message": "%s" }`, inbound.Caller, inbound.Message)))
}
