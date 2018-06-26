package service

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/navono/go-microservice/accountservice/util"
)

// NewRouter Function that returns a pointer to a mux.Router
func NewRouter() *mux.Router {
	// Create an instance of the Gorilla router
	router := mux.NewRouter().StrictSlash(true)

	// Iterate over the routes
	for _, route := range routes {
		var handler http.Handler
		handler = route.HandlerFunc
		handler = util.Logger(handler, route.Name)

		// Attach each route
		router.Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)
	}
	return router
}
