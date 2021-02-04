package auth

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

// Routes group all routes of auth aggregation
func Routes(router *mux.Router) {
	authRouter := router.PathPrefix("/auth").Subrouter()

	authRouter.HandleFunc("", func(resWriter http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(resWriter, "Auth Router")
	}).Methods("POST")
}
