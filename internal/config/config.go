package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type ApplicationConfig struct {
	General  *GeneralConfig
	Database *DatabaseConfig
	Http     *HttpConfig
}

func LoadConfig() *ApplicationConfig {
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatal("Failed to load .env file:", err)
		os.Exit(1)
	}

	cfg := &ApplicationConfig{
		General:  NewGeneralConfig(),
		Http:     NewHttpConfig(),
		Database: NewDatabaseConfig(),
	}

	return cfg
}
