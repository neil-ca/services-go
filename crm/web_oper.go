package main

import (
	"fmt"
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
func Create(w http.ResponseWriter, r *http.Request) {
	template_html.ExecuteTemplate(w, "Create", nil)
}

// Insert - execute template
func Insert(w http.ResponseWriter, r *http.Request) {
	var customer Customer
	customer.CustomerName = r.FormValue("customerName")
	customer.SSN = r.FormValue("ssn")
	InsertCustomer(customer)
	var customers []Customer
	customers = GetCustomers()
	template_html.ExecuteTemplate(w, "Home", customers)
}

// Alter - execute template
func Alter(w http.ResponseWriter, r *http.Request) {
	var customer Customer
	var customerId int
	var customerIdStr string
	customerIdStr = r.FormValue("id")
	fmt.Scanf(customerIdStr, "%d", &customerId)
	customer.CustomerId = customerId
	customer.CustomerName = r.FormValue("ssn")
	UpdateCustomer(customer)
	var customers = GetCustomers()
	template_html.ExecuteTemplate(w, "Home", customers)
}

// Update - execute template
func Update(w http.ResponseWriter, r *http.Request) {
	var customerId int
	var customerIdStr string
	customerIdStr = r.FormValue("id")
	fmt.Sscanf(customerIdStr, "%d", &customerId)
	var customer Customer
	customer = GetCustomerById(customerId)
	template_html.ExecuteTemplate(w, "Update", customer)
}

// Delete -execute Template
func Delete(w http.ResponseWriter, r *http.Request) {
	var customerId int
	var customerIdStr string
	customerIdStr = r.FormValue("id")
	fmt.Sscanf(customerIdStr, "%d", &customerId)
	var customer Customer
	customer = GetCustomerById(customerId)
	DeleteCustomer(customer)
	var customers []Customer
	customers = GetCustomers()
	template_html.ExecuteTemplate(w, "Home", customers)
}

// View - execute Template
func View(w http.ResponseWriter, r *http.Request) {
	var customerId int
	var customerIdStr string
	customerIdStr = r.FormValue("id")
	fmt.Sscanf(customerIdStr, "%d", &customerId)
	var customer Customer
	customer = GetCustomerById(customerId)
	fmt.Println(customer)
	var customers []Customer
	customers = []Customer{customer}
	customers.append(customer)
	template_html.ExecuteTemplate(w, "View", customers)
}
func main() {
	log.Println("Server running on -> http://localhost:8000")
	http.HandleFunc("/", Home)
	http.HandleFunc("/alter", Alter)
	http.HandleFunc("/create", Create)
	http.HandleFunc("/update", Update)
	http.HandleFunc("/view", View)
	http.HandleFunc("/insert", Insert)
	http.HandleFunc("/delete", Delete)

	http.ListenAndServe(":8000", nil)
}
