package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/kcanamar/go-crud/initializers"
	"github.com/kcanamar/go-crud/models"
)

func PostCreate(c *gin.Context) {

	// Get data off request body
	// create a variable for the request body
	var reqBody struct{
		Body string
		Title string
	}
	// parses the request's body as JSON 
	c.Bind(&reqBody)
	 

	// Create a post - https://gorm.io/docs/create.html

	// hardcoded example
	// post := models.Post{Title: "Yes", Body: "We did it!"}

	// request body example
	post := models.Post{Title: reqBody.Title, Body: reqBody.Body}

	result := initializers.DB.Create(&post)
	
	// handle error from result
	if result.Error != nil {
		c.Status(400)
		return
	}

	// return post
	c.JSON(200, gin.H{
		// send back a confirmation of the post
		"post": post,
	})
}