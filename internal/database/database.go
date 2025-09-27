package database

import (
	"context"
	"fmt"
	"log"

	"github.com/matheuscaet/go-api-template/internal/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Connect() *mongo.Client {
	fmt.Println("MongoURI: ", config.MongoURI)
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(config.MongoURI))
	if err != nil {
		log.Fatalf("Failed to connect to MongoDB: %v", err)
	}
	return client
}
