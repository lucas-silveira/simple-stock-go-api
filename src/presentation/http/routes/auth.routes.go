package routes

import (
	"github.com/gorilla/mux"

	"main/src/presentation/http/handlers"
)

// Auth group all routes of auth aggregation
func Auth(router *mux.Router) {
	authRouter := router.PathPrefix("/auth").Subrouter()

	authRouter.HandleFunc("", handlers.PostAuth).Methods("POST")
}
