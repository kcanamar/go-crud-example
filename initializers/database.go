package initializers

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// create a global variable for Database
var DB *gorm.DB

func ConnectDB() {
	// create local scoped error
	var err error
	dsn := os.Getenv("DATABASE_URL")
	fmt.Println(dsn, "< this is dsn")
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	// error handling
	if err != nil {
		log.Fatal("Failed to connect to database")
	}

}