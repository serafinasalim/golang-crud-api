package database

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

var DB *sql.DB

const (
	DBUser     = "user"
	DBPassword = "password"
	DBName     = "dbname"
)

func ConnectDB() {
	connStr := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", DBUser, DBPassword, DBName)
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
