package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/matheuscaet/go-api-template/api/grpc"
	"github.com/matheuscaet/go-api-template/api/handlers"
	consumer "github.com/matheuscaet/go-api-template/consumers"
	"github.com/matheuscaet/go-api-template/internal/config"
)

func main() {
	fmt.Println("Starting Go API")
	config.LoadEnvVariables()

	go func() {
		log.Println("Starting RabbitMQ consumer in goroutine...")
		consumer.Start()
	}()

	go func() {
		log.Println("Starting API server in goroutine...")
		handlers.StartServer()
	}()

	go func() {
		log.Println("Starting gRPC server in goroutine...")
		grpc.StartGRPCServer()
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	<-quit

	log.Println("Shutting down gracefully...")
}
