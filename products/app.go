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

func main() {

}
