package api

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

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

func Handlers() *gin.Engine {
	r := gin.Default()
	r.Use(Cors())
	v1 := r.Group("api/v1/books")
	{
		v1.POST("", PostBook)
		v1.GET("", GetBooks)
		v1.GET(":id", GetBook)
		v1.DELETE(":id", DeleteBook)
		v1.PUT(":id", UpdateBook)
		v1.OPTIONS("", OptionsBook)    //POST
		v1.OPTIONS(":id", OptionsBook) // PUT, DELETE
	}
	return r
}

// Cors set globally a custom middleware
func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Add("Access-Control-Allow-Origin", "*")
		c.Next()
	}
}
