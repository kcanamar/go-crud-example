package main

import (
	"github.com/gin-gonic/gin"
	"github.com/kcanamar/go-crud/controllers"
	"github.com/kcanamar/go-crud/initializers"
)

func init() {
	initializers.LoadEnv()
	initializers.ConnectDB()
}

func main() {
	// setup a gin router
	r := gin.Default()
	
	// Home Route
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello from the other side!",
		})
	})

	// Index Route
	r.GET("/posts", controllers.PostsIndex)

	// Create Route
	r.POST("/posts", controllers.PostCreate)

	// Show Route
	r.GET("/posts/:id", controllers.PostShow)

	// setup our server listener
	r.Run()
}