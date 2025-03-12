package main

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

var db *sql.DB

func initDB() {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatalf("Error loading .env file")
	}

	sqlps := os.Getenv("SQL_CRED")
	var err error

	db, err = sql.Open("mysql", sqlps)
	if err != nil {
		log.Fatal(err)
	}
}
