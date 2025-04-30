package config

import "os"

type GeneralConfig struct {
	SecretKey []byte
}

func NewGeneralConfig() *GeneralConfig {
	return &GeneralConfig{
		SecretKey: []byte(os.Getenv("SECRET_KEY")),
	}
}
