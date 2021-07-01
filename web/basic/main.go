package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", hello)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// Start the web server
	log.Printf("Server listening on port %s", port)
	log.Fatal(http.ListenAndServe(":"+port, mux))
}

func hello(w http.ResponseWriter, r *http.Request) {
	log.Printf("Serving requests: %s", r.URL.Path)
	host, _ := os.Hostname()
	fmt.Fprintf(w, "Hello, gophers 👽\n")
	fmt.Fprintf(w, "Hostname: %s 👽\n", host)
}
