package service

import (
	"github.com/gorilla/mux"
)

// NewRouter Function that returns a pointer to a mux.Router
func NewRouter() *mux.Router {
	// Create an instance of the Gorilla router
	router := mux.NewRouter().StrictSlash(true)

	// Iterate over the routes
	for _, route := range routes {
		// Attach each route
		router.Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(route.HandlerFunc)
	}
	return router
}
