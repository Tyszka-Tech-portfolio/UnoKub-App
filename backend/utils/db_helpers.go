package utils

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

var db *sql.DB

func InitDB() {
	var err error
	db, err = sql.Open("sqlite3", "./database/your-database.db")
	if err != nil {
		log.Fatalf("Failed to open database: %v", err)
	}
}

func GetDB() *sql.DB {
	return db
}
