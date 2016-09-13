package routes

import (
	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	router := mux.NewRouter()
	apiRoute := router.PathPrefix("/api").Subrouter()
	for _, route := range apiRoutes {
		apiRoute.Methods(route.Method).Path(route.Pattern).Handler(route.HandlerFunc)
	}
	for _, route := range webRoutes {
		router.Methods(route.Method).Path(route.Pattern).Handler(route.HandlerFunc)
	}
	return router
}
