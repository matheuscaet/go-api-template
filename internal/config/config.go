package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var (
	RedisHost     = os.Getenv("REDIS_HOST")
	RedisPort     = os.Getenv("REDIS_PORT")
	RedisPassword = os.Getenv("REDIS_PASSWORD")
	RedisDB       = os.Getenv("REDIS_DB")
)

func LoadEnvVariables() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}
}
