package main

import (
	"github.com/gin-gonic/gin"
	"github.com/lucasdsampaio/gin-rest/controllers"
	"github.com/lucasdsampaio/gin-rest/models"
)

func main() {
	r := gin.Default()
	models.ConnectDataBase()

	//r.GET("/", func(c *gin.Context) {
	//	c.JSON(http.StatusOK, gin.H{"data": "hello world"})
	//})
	r.GET("/books", controllers.FindBooks)
	r.POST("/book", controllers.CreateBook)
	r.GET("/book/:id", controllers.FindBook)
	r.PATCH("/book/:id", controllers.UpdateBook)
	r.DELETE("/book/:id", controllers.DeleteBook)

	r.Run()
}
