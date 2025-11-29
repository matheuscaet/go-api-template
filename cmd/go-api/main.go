package main

import (
	"fmt"

	"github.com/matheuscaet/go-api-template/api/handlers"
	consumer "github.com/matheuscaet/go-api-template/consumers"
	"github.com/matheuscaet/go-api-template/internal/config"
)

func main() {
	fmt.Println("Starting Go API")
	config.LoadEnvVariables()
	consumer.Start()
	handlers.StartServer()
}
