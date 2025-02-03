package database

import (
	"context"
	"fmt"
	"log"

	"github.com/Gierdiaz/Vertex-go/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var DB *mongo.Database

func ConnectMongo(cfg *config.Config) error {
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(cfg.MongoURI))
	if err != nil {
		log.Fatalf("Erro ao conectar ao MongoDB: %v", err)
	}

	// client, err := mongo.Connect(context.Background(), options.Client().ApplyURI("mongodb://localhost:27017"))
	// if err != nil {
	// 	log.Fatalf("Erro ao conectar ao MongoDB: %v", err)
	// }

	if err = client.Ping(context.Background(), readpref.Primary()); err != nil {
		log.Fatalf("Erro ao verificar a conexão com o MongoDB: %v", err)
	}

	DB = client.Database("contacts_db")
	fmt.Println("✅ Conectado ao MongoDB")
	return nil
}
