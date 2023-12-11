package outbox

type OutboundRequest struct {
	Caller   string `json:"caller"`
	Message  string `json:"message"`
	Response string `json:"response"`
}
