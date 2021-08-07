package demo

import (
	"database/sql"
	"log"
)

func CreateDatabaseConnection() *sql.DB {
	// var db *sql.DB

	// Move to config file or environment variables
	db, err := sql.Open("postgres", "postgres://data:password@192.168.75.67/data?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	// Ping to database
	err = db.Ping()
	if err != nil {
		log.Fatal("Failed to ping DB: ", err)
	}

	return db
}
