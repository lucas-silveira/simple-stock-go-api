package routes

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

// AuthRoutes group all routes of auth aggregation
func AuthRoutes(router *mux.Router) {
	authRouter := router.PathPrefix("/auth").Subrouter()

	authRouter.HandleFunc("", func(resWriter http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(resWriter, "Auth Router")
	}).Methods("POST")
}
