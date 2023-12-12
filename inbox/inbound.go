package inbox

import "time"

type InboundRequest struct {
	Caller  string    `json:"caller"`
	Message string    `json:"message"`
	Start   time.Time `json:"start"`
}
