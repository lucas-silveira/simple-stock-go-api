package main

import (
	"log"
	"net/http"
	"os"

	"main/src/infra/api"

	"github.com/gorilla/handlers"
)

func main() {
	loggedRoutes := handlers.LoggingHandler(os.Stdout, api.Router)
	log.Fatal(http.ListenAndServe(":3000", loggedRoutes))
}
