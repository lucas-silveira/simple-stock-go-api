package drivers

import (
	"database/sql"
	"fmt"
	"log"
	"main/src/infra/utils/envconfig"
	"os"

	_ "github.com/lib/pq"
)

func connectionToPostgreSQL() *sql.DB {
	var err error

	err = envconfig.Load()

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	connectionString := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("SQL_HOST"),
		os.Getenv("SQL_USER"),
		os.Getenv("SQL_PASSWORD"),
		os.Getenv("SQL_DATABASE"),
	)

	db, err := sql.Open("postgres", connectionString)

	if err != nil {
		log.Fatal(err)
		return nil
	}

	return db
}

// DBConnection is a connection pool for postgres database
var DBConnection = connectionToPostgreSQL()
