package main

import (
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func main() {
	router := httprouter.New()
	// Mapping to methods is possible with HttpRouter
	router.ServeFiles("/static/*filepath",
		http.Dir("/home/neil/static"))
	log.Fatal(http.ListenAndServe(":8080", router))
}
