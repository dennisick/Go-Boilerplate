package config

import (
	"log"
	"os"
	"strconv"
)

type HttpConfig struct {
	Port int
}

func NewHttpConfig() *HttpConfig {
	httpPort, err := strconv.Atoi(os.Getenv("HTTP_PORT"))
	if err != nil {
		log.Fatal("Failed to convert HTTP_PORT to int:", err)
		os.Exit(1)
	}

	return &HttpConfig{
		Port: httpPort,
	}
}
