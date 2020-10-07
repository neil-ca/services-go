package main

import (
	"log"
	"net/http"
	"text/template"
)

// Home method renders the main.html
func Home(w http.ResponseWriter, r *http.Request) {
	var template_html *template.Template
	template_html = template.Must(template.ParseFiles("main.html"))
	template_html.Execute(w, nil)
}
func main() {
	log.Println("Server running on -> http://localhost:8000")
	http.HandleFunc("/", Home)
	http.ListenAndServe(":8000", nil)
}
