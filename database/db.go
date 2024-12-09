package database

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func ConnectDB() {
	connStr := "user=postgres password=12345 dbname=golang-crud sslmode=disable"
	var err error
	DB, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v\n", err)
	}

	if err := DB.Ping(); err != nil {
		log.Fatalf("Database ping failed: %v\n", err)
	}
	log.Println("Database connected!")
}
