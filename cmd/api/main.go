package main

import (
	"log"

	"github.com/dev-roberto-sousa/nuvemhub/internal/api/config"
	"github.com/dev-roberto-sousa/nuvemhub/internal/api/handlers"
)

func main() {
	config.ConnectDatabase()

	router := handlers.SetupRouter()

	log.Println("Server running on port 8080...")
	err := router.Run(":8080")
	if err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
