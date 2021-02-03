package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
)

// Book is a placeholder for book
type Book struct {
	ID     int    `gorm:"primary key;auto increment;not null" form:"id" json:"id"`
	Name   string `gorm:"not null" form:"name" json:"name"`
	Author string `gorm:"not null" form:"author" json:"author"`
}

// InitDb set our database
func InitDb() *gorm.DB {
	// opening file
	db, err := gorm.Open("sqlite3", "./data.db")
	if err != nil {
		panic(err)
	}
	// Display SQL queries
	db.LogMode(true)

	// Creating the table
	if !db.HasTable(&Book{}) {
		db.CreateTable(&Book{})
		// db.Set("gorm:table_options", "ENGINE=InnoDB").CreateTable(&Book{})
	}
	return db
}

func main() {
	// db, err := sql.Open("sqlite3", "./books.db")
	// log.Println(db)
	// if err != nil {
	// 	log.Println(err)
	// }
	// //Create table
	// statement, err := db.Prepare("CREATE TABLE IF NOT EXISTS books (id INTEGER PRIMARY KEY, isbn INTEGER, author VARCHAR(64), name VARCHAR(64) NULL)")
	// if err != nil {
	// 	log.Println("Error in creating table..")
	// } else {
	// 	log.Println("Succesfully created table books!")
	// }
	// statement.Exec()
	// // Create
	// statement, _ = db.Prepare("INSERT INTO books (name, author, isbn) VALUES (?,?,?)")
	// statement.Exec("A table of Two cities", "Charles Dickens", 1548162)
	// log.Println("Inserted the book into database!")
	// // Read
	// rows, _ := db.Query("SELECT id, name, author FROM books")
	// var tempBook Book
	// for rows.Next() {
	// 	rows.Scan(&tempBook.id, &tempBook.name, &tempBook.author)
	// 	log.Printf("ID:%d, Book:%s, Author:%s\n", tempBook.id,
	// 		tempBook.name, tempBook.author)
	// }
	// // Update
	// statement, _ = db.Prepare("update books set name=? where id=?")
	// statement.Exec("the Tale of Two Cities", 1)
	// log.Println("Succesfully updated the book in database!")
	// // Delete
	// statement, _ = db.Prepare("delete from books where id=?")
	// statement.Exec(1)
	// log.Println("Succesfully deleted the book in database!")
}
