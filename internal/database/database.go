package database

import (
	"context"
	"log"
	"strconv"

	"github.com/go-redis/redis/v8"
	"github.com/matheuscaet/go-api-template/internal/config"
)

func Connect() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     config.RedisHost + ":" + config.RedisPort,
		Password: config.RedisPassword,
		DB:       getRedisDB(),
	})

	_, err := client.Ping(context.Background()).Result()
	if err != nil {
		log.Fatalf("Failed to connect to Redis: %v", err)
	}

	return client
}

func getRedisDB() int {
	db, err := strconv.Atoi(config.RedisDB)
	if err != nil {
		log.Printf("Failed to convert RedisDB to int: %v", err)
		db = 0
	}
	return db
}
