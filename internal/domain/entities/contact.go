package entities

import (
	"errors"
	"regexp"

	"github.com/Gierdiaz/Vertex-go/internal/domain/valueobjects"
	"github.com/Gierdiaz/Vertex-go/internal/validation"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Contact struct {
	ID       primitive.ObjectID   `bson:"_id" json:"id"`
	Nome     string               `bson:"nome" json:"nome" validate:"required"`
	Email    string               `bson:"email" json:"email" validate:"required,email"`
	Telefone string               `bson:"telefone" json:"telefone" validate:"required"`
	Address  valueobjects.Address `bson:"address" json:"address"`
}

func NewContact(Nome, Email, Telefone string, Address valueobjects.Address) (*Contact, error) {
	contact := &Contact{
		ID:       primitive.NewObjectID(),
		Nome:     Nome,
		Email:    Email,
		Telefone: Telefone,
		Address:  Address,
	}

	if err := contact.Validate(); err != nil {
		return nil, err
	}

	return contact, nil
}


func (c *Contact) Validate() error {
	if c.Nome == "" {
		return errors.New("nome não pode ser vazio")
	}
	if !isValidEmail(c.Email) {
		return errors.New("e-mail inválido. O formato correto é: exemplo@dominio.com")
	}
	if !isValidTelefone(c.Telefone) {
		return errors.New("telefone inválido. O formato correto é: DDD + número")
	}

	if err := validation.ValidateStruct(c); err != nil {
		return err
	}

	return nil
}

func isValidEmail(email string) bool {
	re := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	return re.MatchString(email)
}

func isValidTelefone(telefone string) bool {
	re := regexp.MustCompile(`^\d{2}\d{8,9}$`)
	return re.MatchString(telefone)
}
