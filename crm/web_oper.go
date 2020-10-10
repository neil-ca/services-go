package main

import (
	"log"
	"net/http"
	"text/template"
)

var template_html = template.Must(template.ParseGlob("templates/*"))

// Home - execute Template
func Home(w http.ResponseWriter, r *http.Request) {
	var customers []Customer
	customers = GetCustomers()
	log.Println(customers)
	template_html.ExecuteTemplate(w, "Home", customers)
}

// Create - execute template
func main() {
	log.Println("Server running on -> http://localhost:8000")
	http.HandleFunc("/", Home)
	http.ListenAndServe(":8000", nil)
}
