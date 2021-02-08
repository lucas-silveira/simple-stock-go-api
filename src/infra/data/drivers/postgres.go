package drivers

import (
	"database/sql"
	"fmt"
	"log"
	"os"
)

func connectionToPostgreSQL() *sql.DB {
	connectionString := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s",
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
