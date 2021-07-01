package main

import (
	"log"
	"net/http"

	"github.com/Neil-uli/Restful-go/books-api/api"
	_ "github.com/mattn/go-sqlite3"
)

func main() {

	err := http.ListenAndServe(":8080", api.Handlers())
	//err := http.ListenAndServeTLS(":3000", "certificate/localhost.pem", "certificate/localhost.key", api.Handlers())

	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}

}
