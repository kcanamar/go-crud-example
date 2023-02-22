package main

import (
	"github.com/gin-gonic/gin"
	"github.com/kcanamar/go-crud/initializers"
)

func init() {
	initializers.LoadEnv()
}

func main() {
	// setup a gin router
	r := gin.Default()
	// Declare home get route
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello from the other side!",
		})
	})
	// setup our server listener
	r.Run()
}