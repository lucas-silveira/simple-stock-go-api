package main

import (
	auth "main/src/domain/auth"

	"github.com/gorilla/mux"
)

var routes = func() *mux.Router {
	router := mux.NewRouter()

	auth.Routes(router)

	return router
}()
