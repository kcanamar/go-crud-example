package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/kcanamar/go-crud/initializers"
	"github.com/kcanamar/go-crud/models"
)

func PostCreate(c *gin.Context) {

	// Get data off request body
	// https://gin-gonic.com/docs/examples/binding-and-validation/
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
	newPost := models.Post{Title: reqBody.Title, Body: reqBody.Body}

	// create post in database with newPost 
	result := initializers.DB.Create(&newPost)
	
	// handle error from result
	if result.Error != nil {
		c.Status(400)
		return
	}

	// return post
	c.JSON(200, gin.H{
		// send back a confirmation of the post
		"post": newPost,
	})
}

func PostsIndex(c *gin.Context) {
	// Get all posts
	// https://gorm.io/docs/query.html#Retrieving-all-objects

	// create an array typed models.Post to hold all posts
	var allPosts []models.Post

	// query the database
	initializers.DB.Find(&allPosts)

	// Response
	c.JSON(200, gin.H{
		"posts": allPosts,
	})
}

func PostShow(c *gin.Context) {
	// Get single post by id
	// https://gorm.io/docs/query.html#Retrieving-objects-with-primary-key

	// variable to hold all post typed models.Post
	var post models.Post

	// query the database
	initializers.DB.First(&post, c.Param("id"))

	// Response
	c.JSON(200, gin.H{
		"post": post,
	})
}

func PostUpdate(c *gin.Context){
	// Get the id from params
	id := c.Param("id")

	// Get the request body and bind to context
	var body struct{
		Title string
		Body string
	}
	c.Bind(&body)

	// Find post to be updated
	var post models.Post
	initializers.DB.Find(&post, id)

	// https://gorm.io/docs/update.html#Updates-multiple-columns
	// Update Post
	initializers.DB.Model(&post).Updates(models.Post{
		Title: body.Title, Body: body.Body,
	})

	// Response
	c.JSON(200, gin.H{
		"updatedPost": post,
	})
}

func PostDestroy(c *gin.Context) {
	// Get id from params
	id := c.Param("id")

	// https://gorm.io/docs/delete.html#Delete-with-primary-key
	// Delete the post
	initializers.DB.Delete(&models.Post{}, id)
	
	// Response
	c.JSON(200, gin.H{
		"message": "Successful Destory",
	})
}