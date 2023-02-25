package initializers

import (
	"log"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
)

func LoadEnv() {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	environmentPath := filepath.Join(dir, ".env")
	err = godotenv.Load(environmentPath)
	
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}