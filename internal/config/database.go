package config

import (
	"log"
	"os"
	"strconv"
)

type DatabaseConfig struct {
	Host     string
	Port     int
	User     string
	Password string
	Database string
}

func NewDatabaseConfig() *DatabaseConfig {
	databasePort, err := strconv.Atoi(os.Getenv("POSTGRES_PORT"))
	if err != nil {
		log.Fatal("Failed to convert POSTGRES_PORT to int:", err)
		os.Exit(1)
	}

	return &DatabaseConfig{
		Host:     os.Getenv("POSTGRES_HOST"),
		Port:     databasePort,
		User:     os.Getenv("POSTGRES_USER"),
		Password: os.Getenv("POSTGRES_PASSWORD"),
		Database: os.Getenv("POSTGRES_DATABASE"),
	}
}
