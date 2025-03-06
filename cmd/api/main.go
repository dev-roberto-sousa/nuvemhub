package main

import (
	"log"

	"github.com/dev-roberto-sousa/nuvemhub/internal/api/handlers"
)

func main() {
	router := handlers.SetupRouter()

	log.Println("Servidor rodando na porta 8080...")
	err := router.Run(":8080")
	if err != nil {
		log.Fatalf("Erro ao iniciar o servidor: %v", err)
	}
}
