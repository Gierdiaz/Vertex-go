package interfaces

import (
	"github.com/Gierdiaz/Vertex-go/internal/entities"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ContactRepositoryInterface interface {
	Create(contact *entities.Contact) error
	GetAll() ([]*entities.Contact, error)
	FindById(id primitive.ObjectID) (*entities.Contact, error)
}
