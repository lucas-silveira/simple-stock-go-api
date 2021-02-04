package routes

import (
	"github.com/gorilla/mux"

	"main/src/infra/api/handlers"
)

// Auth group all routes of auth aggregation
func Auth(router *mux.Router) {
	authRouter := router.PathPrefix("/auth").Subrouter()

	authRouter.HandleFunc("", handlers.PostAuth).Methods("POST")
}
