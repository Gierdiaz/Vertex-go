package handlers

import (
	"log"

	"github.com/Gierdiaz/Vertex-go/internal/application/services"
	"github.com/Gierdiaz/Vertex-go/internal/domain/entities"
	"github.com/Gierdiaz/Vertex-go/internal/domain/valueobjects"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ContactHandler struct {
	service *services.ContactService
}

func NewContactHandler(service *services.ContactService) *ContactHandler {
	return &ContactHandler{service: service}
}

func (h *ContactHandler) CreateContact(c *fiber.Ctx) error {
	var req struct {
		Nome     string `json:"nome"`
		Email    string `json:"email"`
		Telefone string `json:"telefone"`
		Cep      string `json:"cep"`
	}

	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request"})
	}

	contact := &entities.Contact{
		Nome:     req.Nome,
		Email:    req.Email,
		Telefone: req.Telefone,
		Address: valueobjects.Address{
			CEP: req.Cep,
		},
	}

	if err := contact.Validate(); err != nil {
		log.Printf("Validation failed: %v", err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	newContact, err := h.service.CreateContact(contact)
	if err != nil {
		log.Printf("Error creating contact: %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Error creating contact"})
	}

	return c.JSON(fiber.Map{
		"message": "Contato criado com sucesso",
		"contact": newContact,
	})
}

func (h *ContactHandler) GetContacts(c *fiber.Ctx) error {
	contacts, err := h.service.GetAllContacts()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Error getting contacts"})
	}

	return c.JSON(fiber.Map{"contacts": contacts})
}

func (h *ContactHandler) GetContact(c *fiber.Ctx) error {
	idStr := c.Query("id")
	ID, err := primitive.ObjectIDFromHex(idStr)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid ID"})
	}

	contact, err := h.service.GetContact(ID)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Contato não encontrado"})
	}

	return c.JSON(fiber.Map{"contact": contact})
}
