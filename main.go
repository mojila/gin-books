package main

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/mojila/gin-books/controllers"
	"github.com/mojila/gin-books/models"
)

func main() {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"data": "hello world"})
	})

	r.GET("/books", controllers.FindBooks)
	r.GET("/books/:id", controllers.FindBook)
	r.POST("/books", controllers.CreateBooks)
	r.PATCH("/books/:id", controllers.UpdateBook)
	r.DELETE("/books/:id", controllers.DeleteBook)

	models.ConnectDatabase()

	r.Run()
}
