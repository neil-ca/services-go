package main

import (
	"database/sql"
	"log"

	"github.com/Neil-Uli/rest/metroRailAPI/dbutils"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	// Connect to database
	db, err := sql.Open("sqlite3", "./railapi.db")
	if err != nil {
		log.Println("Driver creation failed!")
	}
	// Create tables
	dbutils.Initialize(db)
}
