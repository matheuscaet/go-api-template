package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var (
	MongoURI = ""
	Port     = ""
)

func LoadEnvVariables() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}
	MongoURI = os.Getenv("MONGO_URI")
	Port = os.Getenv("PORT")
	if Port == "" {
		Port = "8080"
	}
}
