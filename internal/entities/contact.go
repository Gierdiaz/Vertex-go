package entities

import (
	"errors"
	"regexp"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Contact struct {
	ID       primitive.ObjectID `bson:"_id" json:"id"`
	Nome     string             `bson:"nome" json:"nome"`
	Email    string             `bson:"email" json:"email"`
	Telefone string             `bson:"telefone" json:"telefone"`
	Address  Address            `bson:"address_id" json:"address"`
}

func NewContact(Nome, Email, Telefone string, Address Address) (*Contact, error) {
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
		return errors.New("e-mail inválido")
	}
	if !isValidTelefone(c.Telefone) {
		return errors.New("telefone inválido")
	}

	return nil
}

func isValidEmail(email string) bool {
	re := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	return re.MatchString(email)
}

func isValidTelefone(telefone string) bool {
	re := regexp.MustCompile(`^\(?\d{2}\)? ?\d{4,5}-?\d{4}$`)
	return re.MatchString(telefone)
}
