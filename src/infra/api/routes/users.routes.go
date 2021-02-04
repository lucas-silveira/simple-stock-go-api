package routes

import (
	"github.com/gorilla/mux"

	"main/src/infra/api/handlers"
	. "main/src/infra/api/middlewares"
)

// Users group all routes of auth aggregation
func Users(router *mux.Router) {
	authRouter := router.PathPrefix("/users").Subrouter()

	authRouter.Use(AuthMiddleware)

	authRouter.HandleFunc("", handlers.GetAllUsers).Methods("GET")
}
