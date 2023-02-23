package main

import (
	"github.com/kcanamar/go-crud/initializers"
	"github.com/kcanamar/go-crud/models"
)

func init() {
	// import initializers
	initializers.LoadEnv()
	initializers.ConnectDB()
}

func main() {
	// automigrate to database
	initializers.DB.AutoMigrate(&models.Post{})
}