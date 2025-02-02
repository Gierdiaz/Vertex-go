package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	MongoURI string
	DBName   string
}

func Load() (*Config, error) {
	if err := godotenv.Load("/app/.env"); err != nil {
		return nil, fmt.Errorf("error loading .env file: %v", err)
	}

	config := &Config{
		MongoURI: os.Getenv("MONGODB_URI"),
		DBName:   os.Getenv("DB_NAME"),
	}

	if config.DBName == "" || config.MongoURI == "" {
		log.Fatalf("As variáveis de ambiente MONGO_URI ou DB_NAME não foram definidas")
	}

	return config, nil
}
