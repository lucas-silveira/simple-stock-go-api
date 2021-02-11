package main

import (
	"log"
	"net/http"
	"os"

	"main/src/infra/utils/envconfig"
	. "main/src/presentation/http"

	"github.com/gorilla/handlers"
)

func main() {
	err := envconfig.Load()

	if err != nil {
		log.Fatalf("Error loading .env file, err: %s", err)
	}

	loggedRoutes := handlers.LoggingHandler(os.Stdout, Router)
	log.Fatal(http.ListenAndServe(":3000", loggedRoutes))
}
