package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func Init() error {
	var err error
	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("POSTGRES_HOST"),
		os.Getenv("POSTGRES_PORT"),
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"),
		os.Getenv("POSTGRES_DATABASE"),
	)
	// Inisialisasi variabel global DB
	DB, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Println("Failed to connect to database:", err)
		return err
	}

	if err = DB.Ping(); err != nil {
		log.Println("Database is not reachable:", err)
		return err
	}

	log.Println("Database connected successfully!")
	return nil
}
