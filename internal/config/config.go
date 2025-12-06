package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var (
	MongoURI             = ""
	RabbitMQURI          = ""
	RabbitMQExchange     = "TASKS_EXCHANGE"
	RabbitMQExchangeType = "topic"
	RabbitMQQueue        = "TASKS_QUEUE"
	RabbitMQRoutingKey   = "task.create"
	Port                 = "8080"
)

func LoadEnvVariables() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}
	MongoURI = os.Getenv("MONGO_URI")
	RabbitMQURI = os.Getenv("RABBITMQ_URI")
	RabbitMQExchange = os.Getenv("RABBITMQ_EXCHANGE")
	RabbitMQExchangeType = os.Getenv("RABBITMQ_EXCHANGE_TYPE")
	RabbitMQQueue = os.Getenv("RABBITMQ_QUEUE")
	RabbitMQRoutingKey = os.Getenv("RABBITMQ_ROUTING_KEY")
	Port = os.Getenv("PORT")
}
