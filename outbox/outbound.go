package outbox

import "time"

type OutboundRequest struct {
	Caller   string    `json:"caller"`
	Message  string    `json:"message"`
	Response string    `json:"response"`
	Start    time.Time `json:"start"`
	End      time.Time `json:"end"`
}
