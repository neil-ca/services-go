package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

// Customer class
type Customer struct {
	CustomerId   int
	CustomerName string
	SSN          string
}

// GetConnection method which return sql.DB
func GetConnection() (database *sql.DB) {
	databaseDriver := "mysql"
	databaseUser := "user"
	databasePass := "secret"
	databaseName := "crm"
	database, error := sql.Open(databaseDriver, databaseUser+":"+databasePass+"@/"+databaseName)
	if error != nil {
		panic(error.Error())
	}
	return database
}

// GetCustomers method returns Customer Array
func GetCustomers() []Customer {
	var database *sql.DB
	database = GetConnection()

	var error error
	var rows *sql.Rows
	rows, error = database.Query("SELECT * FROM Customer ORDER BY Customerid DESC")
	if error != nil {
		panic(error.Error())
	}

	var customer Customer
	customer = Customer{}

	var customers []Customer
	customers = []Customer{}
	for rows.Next() {
		var customerId int
		var customerName string
		var ssn string
		error = rows.Scan(&customerId, &customerName, &ssn)
		if error != nil {
			panic(error.Error())
		}
		customer.CustomerId = customerId
		customer.CustomerName = customerName
		customer.SSN = ssn
		customers = append(customers, customer)
	}
	defer database.Close()
	return customers
}

// InsertCustomer method with parameter customer
func InsertCustomer(customer Customer) {
	var database *sql.DB
	database = GetConnection()

	var error error
	var insert *sql.Stmt
	insert, error = database.Prepare("INSERT INTO customer(CustomerName, ssn) VALUES(?,?)")
	if error != nil {
		panic(error.Error())
	}
	insert.Exec(customer.CustomerName, customer.SSN)
	defer database.Close()
}
func main() {
	var customers []Customer
	customers = GetCustomers()
	fmt.Println("Before Insert", customers)
	var customer Customer
	customer.CustomerName = "Richard Stallman"
	customer.SSN = "5415475"
	InsertCustomer(customer)
	customers = GetCustomers()
	fmt.Println("Customers After insert", customers)
}

/*
CREATE TABLE customer (
     customerId INT(10) UNSIGNED AUTO_INCREMENT PRIMARY KEY,
     customerName VARCHAR(30) NOT NULL,
     ssn VARCHAR(30) NOT NULL
	 );
*/
