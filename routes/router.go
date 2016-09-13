package routes

import (
	"github.com/gorilla/mux"
)

func ApiRouter() *mux.Router {
	router := mux.NewRouter()
	apiRoute := router.PathPrefix("/api").Subrouter()
	for _, route := range routes {
		apiRoute.
			Methods(route.Method).
			Path(route.Pattern).
			Handler(route.HandlerFunc)
	}
	return apiRoute
}
