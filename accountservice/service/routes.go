package service

import (
	"net/http"
)

// Route define a single route
type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

// Routes route list
type Routes []Route

var routes = Routes{
	Route{
		"GetAccount",
		"Get",
		"/accounts/{accountId}",
		func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application.json; charset=UTF-8")
			w.Write([]byte("{\"result\":\"OK\"}\n"))
		},
	},
}
