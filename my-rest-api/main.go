package main

// go get -u github.com/gin-gonic/gin
// go mod tidy (for updating dependencies)
import (
	"fmt"
	"my-rest-api/structures"
	"my-rest-api/controllers"
	"net/http" // For HTTP status codes and server functionalities
	"github.com/gin-gonic/gin" // Gin framework
)

// := declares a new variable and assigns it a value at the same time
func main() {
	r := gin.Default()
	book := structures.Book{
		ID:     4,
		Title:  "book title",
		Author: "book author",
	}

	// pritning something
	fmt.Printf("book author %v\n", book.Author);

	r.GET("/hello", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello, World!",
		})
	})

	// function name should be
	r.GET("/get-books", controllers.GetBooks)
	r.POST("/add-books", )

	r.Run(":8080") 
}

