package main

import (
	"log"
	"net/http"
	"os"

	"main/src/infra/api"

	"github.com/gorilla/handlers"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load("development.env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	loggedRoutes := handlers.LoggingHandler(os.Stdout, api.Router)
	log.Fatal(http.ListenAndServe(":3000", loggedRoutes))
}
