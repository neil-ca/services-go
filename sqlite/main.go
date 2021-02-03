package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
)

// Book is a placeholder for book
type Book struct {
	// gorm by default take ID fiel as the table's PK
	ID     int    `gorm:"AUTO_INCREMENT" form:"id" json:"id"`
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

// Cors set
func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Add("Access-Control-Allow-Origin", "*")
		c.Next()
	}
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
	r := gin.Default()
	r.Use(Cors())
	v1 := r.Group("api/v1")
	{
		v1.POST("/book", PostBook)
		v1.GET("/books", GetBooks)
	}
	r.Run(":8080")
}

// PostBook handle the creation of a book
func PostBook(c *gin.Context) {
	db := InitDb()
	defer db.Close()

	var book Book
	// Bind checks the Content-Type to select a binding engine automatically
	c.Bind(&book)

	if book.Name != "" && book.Author != "" {
		// INSERT INTO "books" ("name","author") VALUES ('1984','George Orwell')
		db.Create(&book)
		c.JSON(201, gin.H{"success": book})
	} else {
		// Display error
		c.JSON(422, gin.H{"error": "Fields are empty"})
	}
	// curl -i -X POST -H "Content-Type: application/json" -d "{ \"name\": \"1984\", \"author\": \"George Orwell\" }" http://localhost:8080/api/v1/book
}

//GetBooks returns all records and return a JSON response
func GetBooks(c *gin.Context) {
	db := InitDb()
	defer db.Close()

	var books []Book
	c.Bind(&books)
	// SELECT * FROM books;
	db.Find(&books)

	// Display JSON result
	c.JSON(200, books)
	// curl -i http://localhost:8080/api/v1/books
}

func GetBook(c *gin.Context) {
	db := InitDb()
	defer db.Close()

	id := c.Params.ByName("id")
	var book Book
	// SELECT * FROM books WHERE id = ?
	db.First(&book, id)
	if book.ID != 0 {
		c.JSON(200, book)
	} else {
		c.JSON(404, "Book not found")
	}
}
