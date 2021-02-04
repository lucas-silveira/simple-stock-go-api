package routes

import (
	"github.com/gorilla/mux"

	"main/src/infra/api/handlers"
)

// AuthRoutes group all routes of auth aggregation
func AuthRoutes(router *mux.Router) {
	authRouter := router.PathPrefix("/auth").Subrouter()

	authRouter.HandleFunc("", handlers.PostAuth).Methods("POST")
}
