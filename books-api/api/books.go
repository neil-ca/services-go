package api

import (
	"github.com/gin-gonic/gin"
)

// Book is a placeholder for book
type Book struct {
	// gorm by default take ID field as the table's PK
	ID          int    `gorm:"AUTO_INCREMENT" form:"id" json:"id"`
	Name        string `gorm:"not null" form:"name" json:"name"`
	Author      string `gorm:"not null" form:"author" json:"author"`
	Description string `form:"description" json:"description"`
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
		c.JSON(404, gin.H{"error": "Book not found"})
	}
}

func DeleteBook(c *gin.Context) {
	db := InitDb()
	defer db.Close()
	// Param is a simple shorcut of Params.ByName("id")
	id := c.Param("id")
	var book Book

	db.First(&book, id)

	// 204 is recomended if nothing is returned
	//c.JSON(204, book)

	if book.ID != 0 {
		db.Delete(&book)
		c.JSON(200, gin.H{"success": "Book #" + id + " deleted"})
	} else {
		c.JSON(404, gin.H{"error": "Book not found"})
	}
	// curl -i -X DELETE http://localhost:8080/api/v1/book/1
}
func UpdateBook(c *gin.Context) {
	db := InitDb()
	defer db.Close()

	id := c.Param("id")
	var book Book

	db.First(&book, id)

	if book.Name != "" && book.Author != "" {
		if book.ID != 0 {
			var newBook Book
			c.Bind(&newBook)

			result := Book{
				ID:     book.ID,
				Name:   newBook.Name,
				Author: newBook.Author,
			}
			// UPDATE books SET name='newBook.name', author='newBook.author' WHERE id = book.ID;
			db.Save(&result)
			// Display modified data in JSON message "success"
			c.JSON(200, gin.H{"success": result})
		} else {
			c.JSON(404, "Book not found")
		}
	} else {
		c.JSON(422, "Fields are empty")
	}
	// curl -i -X PUT -H "Content-Type: application/json" -d "{ \"name\": \"Un mundo feliz\", \"Author\": \"Aldous Huxley\" }" http://localhost:8080/api/v1/book/1
}
