package main

import (
	"log"

	"github.com/Gierdiaz/Vertex-go/config"
	"github.com/Gierdiaz/Vertex-go/router"
	"github.com/gofiber/fiber/v2"
)

func main() {
	if err := config.ConnectMongo(); err != nil {
		log.Fatalf("Erro ao conectar ao MongoDB: %v", err)
	}

	app := fiber.New()
	router.SetupRoutes(app)

	log.Println("🚀 Servidor rodando em http://localhost:3000")
	if err := app.Listen(":3000"); err != nil {
		log.Fatalf("Erro ao iniciar o servidor: %v", err)
	}
}
