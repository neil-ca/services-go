package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strings"
)

func sayHelloName(w http.ResponseWriter, r *http.Request) {
	r.ParseForm() // Parse url parameters passed, then parse the response packet for
	// the POST body(request body) if you do not call ParseForm, the following data can not be obtained form
	fmt.Println(r.Form)
	fmt.Println("scheme", r.URL.Path)
	fmt.Println("scheme", r.URL.Scheme)
	fmt.Println(r.Form["url_long"])
	for k, v := range r.Form {
		fmt.Println("key: ", k)
		fmt.Println("val: ", strings.Join(v, ""))
	}
	fmt.Fprintf(w, "Hello neil")
}

func login(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method", r.Method) // get request method
	if r.Method == "GET" {
		t, _ := template.ParseFiles("login.gtpl")
		t.Execute(w, nil)
	} else {
		r.ParseForm()
		// logic part of log in
		fmt.Println("username", r.FormValue("username"))
		fmt.Println("password", r.Form["password"])
	}
}

func main() {
	http.HandleFunc("/", sayHelloName)
	http.HandleFunc("/login", login)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("Error listening on 8080", err)
	}
}
