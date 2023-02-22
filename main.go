package main

import (
	"github.com/gin-gonic/gin"
)

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