package main

import (
	"fmt"

	"github.com/matheuscaet/go-api-template/api/handlers"
)

func main() {
	fmt.Println("Starting Go API")
	handlers.StartServer()
}
