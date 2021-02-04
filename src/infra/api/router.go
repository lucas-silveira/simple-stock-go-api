package api

import (
	"github.com/gorilla/mux"

	"main/src/infra/api/middlewares"
	"main/src/infra/api/routes"
)

// Router is mux Router struct that group all app routes
var Router = func() *mux.Router {
	router := mux.NewRouter()

	router.Use(middlewares.AuthMiddleware)
	routes.AuthRoutes(router)

	return router
}()
