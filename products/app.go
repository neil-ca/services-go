package main

import (
	"database/sql"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

// App exposes references to the router and the database
type App struct {
	Router *mux.Router
	DB     *sql.DB
}

// Initialize will take in the details required to connect to the database It will create
// a database connection and wire up the routes to respond according to the requirements
func (a *App) Initialize(user, password, dbname string) {}

// Run method will simply start the application
func (a *App) Run(addr string) {}
