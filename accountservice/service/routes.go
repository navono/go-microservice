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
		"GetAccount",            // Name
		"GET",                   // HTTP method
		"/accounts/{accountID}", // Route pattern
		GetAccount,
	},
	Route{
		"HealthCheck",
		"GET",
		"/health",
		HealthCheck,
	},
}
