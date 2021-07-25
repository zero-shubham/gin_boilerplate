package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DbUri      string
	ServerPort string
}

func LoadConfig() (*Config, error) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
		return nil, err
	}

	return &Config{
		DbUri:      os.Getenv("DB_URI"),
		ServerPort: os.Getenv("SERVER_PORT"),
	}, nil
}
