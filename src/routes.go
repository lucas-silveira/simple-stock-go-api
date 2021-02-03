package main

import (
	authPresentation "main/src/auth/presentation"

	"github.com/gorilla/mux"
)

var routes = func() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)

	authPresentation.Routes(router)

	return router
}()
