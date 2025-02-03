package main

import (
	"log"

	"github.com/Gierdiaz/Vertex-go/config"
	"github.com/Gierdiaz/Vertex-go/internal/infrastructure/database"
	"github.com/Gierdiaz/Vertex-go/internal/interfaces/http/router"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func main() {

	cfg, err := config.LoadEnv()
	if err != nil {
		log.Fatalf("Erro ao carregar configurações do arquivo .env: %v", err)
		return
	}

	if err := database.ConnectMongo(cfg); err != nil {
		log.Fatalf("Erro ao conectar ao MongoDB: %v", err)
	}

	app := fiber.New()

	app.Use(logger.New())  // Middleware de logs
	app.Use(recover.New()) // Middleware de recuperação

	router.SetupRoutes(app)

	log.Println("🚀 Servidor rodando em http://localhost:3000")
	if err := app.Listen(":3000"); err != nil {
		log.Fatalf("Erro ao iniciar o servidor: %v", err)
	}
}
