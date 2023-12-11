package ping

import (
	"encoding/json"
	"net/http"
)

const (
	PONG = "pong"
)

func PingRouteHandler(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(PONG)
}
