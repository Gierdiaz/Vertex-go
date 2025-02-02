package contracts

import (
	"github.com/Gierdiaz/Vertex-go/internal/domain/entities"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ContactRepositoryInterface interface {
	Create(contact *entities.Contact) error
	GetAll() ([]*entities.Contact, error)
	FindById(id primitive.ObjectID) (*entities.Contact, error)
}
