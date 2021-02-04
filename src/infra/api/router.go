package api

import (
	"github.com/gorilla/mux"

	"main/src/infra/api/routes"
)

// Router is mux Router struct that group all app routes
var Router = func() *mux.Router {
	router := mux.NewRouter()

	routes.Auth(router)
	routes.Users(router)

	return router
}()
