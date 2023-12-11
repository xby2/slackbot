package router

import (
	"net/http"

	inbox "slackbot/router/inbox"
	ping "slackbot/router/ping"
)

type Route struct {
	Name        string
	Method      string
	Path        string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{
	Route{
		"Ping",
		"GET",
		"/ping",
		ping.PingRouteHandler,
	},
	Route{
		"Inbox",
		"POST",
		"/inbox",
		inbox.InboxRouteHandler,
	},
}
