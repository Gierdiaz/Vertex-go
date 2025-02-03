package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	MongoURI string
	DBName   string
}

func LoadEnv() (*Config, error) {
	wd, err := os.Getwd()
	if err != nil {
		log.Fatalf("Erro ao obter o diretório de trabalho: %v", err)
	}
	log.Printf("Diretório de trabalho: %s", wd)
	
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Erro ao carregar o arquivo .env: %v", err)
		return nil, err
	}

	config := &Config{
		MongoURI: os.Getenv("MONGODB_URI"),
		DBName:   os.Getenv("DB_NAME"),
	}

	if config.DBName == "" || config.MongoURI == "" {
		log.Fatalf("As variáveis de ambiente MONGODB_URI ou DB_NAME não foram definidas")
	}

	return config, nil
}