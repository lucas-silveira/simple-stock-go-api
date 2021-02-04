package routes

import (
	"github.com/gorilla/mux"

	"main/src/infra/api/handlers"
	"main/src/infra/api/middlewares"
)

// UsersRoutes group all routes of auth aggregation
func UsersRoutes(router *mux.Router) {
	authRouter := router.PathPrefix("/users").Subrouter()

	authRouter.Use(middlewares.AuthMiddleware)

	authRouter.HandleFunc("", handlers.GetAllUsers).Methods("GET")
}
