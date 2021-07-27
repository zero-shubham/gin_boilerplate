package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DbUri      string
	ServerPort string
	SigningKey string
}

var config *Config

func LoadConfig() (*Config, error) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
		return nil, err
	}
	if config == nil {
		config = &Config{
			DbUri:      os.Getenv("DB_URI"),
			ServerPort: os.Getenv("SERVER_PORT"),
			SigningKey: os.Getenv("SIGNING_KEY"),
		}
	}
	return config, nil
}

func GetConfig() *Config {
	if config == nil {
		return nil
	}
	return config
}
