package services

import (
	"github.com/Gierdiaz/Vertex-go/internal/entities"
	"github.com/Gierdiaz/Vertex-go/internal/integrations"
	"github.com/Gierdiaz/Vertex-go/internal/interfaces"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ContactService struct {
	contactRepo interfaces.ContactRepositoryInterface
}

func NewContactService(contactRepo interfaces.ContactRepositoryInterface) *ContactService {
	return &ContactService{contactRepo: contactRepo}
}

func (s *ContactService) CreateContact(contact *entities.Contact) (*entities.Contact, error) {
	addressData, err := integrations.GetAddressByCEP(contact.Address.CEP)
	if err != nil {
		return nil, err
	}

	address, err := entities.NewAddress(
		addressData.CEP,
		addressData.Logradouro,
		addressData.Bairro,
		addressData.Localidade,
		addressData.UF,
		addressData.IBGE,
		addressData.DDD,
		addressData.Siafi,
	)
	if err != nil {
		return nil, err
	}

	newContact, err := entities.NewContact(contact.Nome, contact.Email, contact.Telefone, *address)
	if err != nil {
		return nil, err
	}

	if err := s.contactRepo.Create(newContact); err != nil {
		return nil, err
	}

	return newContact, nil
}

func (s *ContactService) GetAllContacts() ([]*entities.Contact, error) {
	return s.contactRepo.GetAll()
}

func (s *ContactService) GetContact(id primitive.ObjectID) (*entities.Contact, error) {
	return s.contactRepo.FindById(id)
}