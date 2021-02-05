package routes

import (
	"github.com/gorilla/mux"

	"main/src/presentation/http/handlers"
	. "main/src/presentation/http/middlewares"
)

// Users group all routes of auth aggregation
func Users(router *mux.Router) {
	authRouter := router.PathPrefix("/users").Subrouter()

	authRouter.Use(AuthMiddleware)

	authRouter.HandleFunc("", handlers.GetUsers).Methods("GET")
	authRouter.HandleFunc("", handlers.PostUsers).Methods("POST")
}
