package main

import (
	"fmt"

	"github.com/matheuscaet/go-api-template/api/handlers"
	"github.com/matheuscaet/go-api-template/internal/config"
)

func main() {
	fmt.Println("Starting Go API")
	config.LoadEnvVariables()
	handlers.StartServer()
}
