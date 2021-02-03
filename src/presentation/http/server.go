package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
)

func main() {
	loggedRoutes := handlers.LoggingHandler(os.Stdout, routes)
	log.Fatal(http.ListenAndServe(":3000", loggedRoutes))
}
