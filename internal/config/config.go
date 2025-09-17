package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var (
	MongoURI = os.Getenv("MONGO_URI")
)

func LoadEnvVariables() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}
}
