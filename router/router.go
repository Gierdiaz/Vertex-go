package router

import (
	"github.com/Gierdiaz/Vertex-go/config"
	"github.com/Gierdiaz/Vertex-go/internal/handlers"
	"github.com/Gierdiaz/Vertex-go/internal/repositories"
	"github.com/Gierdiaz/Vertex-go/internal/services"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	contactRepo := repositories.NewContactRepository(config.DB)
	contactService := services.NewContactService(contactRepo)
	contactHandler := handlers.NewContactHandler(contactService)

	app.Post("/contacts", contactHandler.CreateContact)
	app.Get("/contacts", contactHandler.GetContacts)
	app.Get("/contacts/:id", contactHandler.GetContact)

	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"message": "🚀 API rodando com Fiber!"})
	})
}
