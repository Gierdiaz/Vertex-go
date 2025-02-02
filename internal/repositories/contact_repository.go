package repositories

import (
	"context"

	"github.com/Gierdiaz/Vertex-go/internal/entities"
	"github.com/Gierdiaz/Vertex-go/internal/interfaces"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type ContactRepository struct {
	collection *mongo.Collection
}

func NewContactRepository(db *mongo.Database) interfaces.ContactRepositoryInterface {
	return &ContactRepository{
		collection: db.Collection("contacts"),
	}
}

func (r *ContactRepository) Create(contact *entities.Contact) error {
	_, err := r.collection.InsertOne(context.Background(), contact)
	return err
}

func (r *ContactRepository) GetAll() ([]*entities.Contact, error) {
	var contacts []*entities.Contact
	cursor, err := r.collection.Find(context.Background(), bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())
	for cursor.Next(context.Background()) {
		var contact entities.Contact
		if err := cursor.Decode(&contact); err != nil {
			return nil, err
		}
		contacts = append(contacts, &contact)
	}
	return contacts, nil
}

func (r *ContactRepository) FindById(id primitive.ObjectID) (*entities.Contact, error) {
	var contact entities.Contact
	err := r.collection.FindOne(context.Background(), bson.M{"_id": id}).Decode(&contact)
	if err != nil {
		return nil, err
	}
	return &contact, nil
}