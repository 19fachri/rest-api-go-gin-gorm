package main

import (
	"example/web-service-go/controllers"
	"example/web-service-go/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	models.ConnectDatabase()

	router.GET("/", func(c *gin.Context) {
		c.IndentedJSON(http.StatusOK, gin.H{"message": "hello world"})
	})
	router.GET("/books", controllers.FindAllBooks)
	router.POST("/books", controllers.CreateBook)
	router.GET("/books/:id", controllers.FindBookById)
	router.PUT("/books/:id", controllers.UpdateBookById)
	router.DELETE("/books/:id", controllers.DeleteBook)

	router.Run("localhost:8080")
}
