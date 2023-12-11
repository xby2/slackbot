package inbox

type InboundRequest struct {
	Caller  string `json:"caller"`
	Message string `json:"message"`
}
