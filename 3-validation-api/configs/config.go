package configs

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

type Config struct {
	EmailConf EmailConfig
}
type EmailConfig struct {
	Email    string
	Password string
	Address  string
}

func LoadConfig() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file, using default config.")
	}
	return &Config{
		EmailConf: EmailConfig{
			Email:    os.Getenv("EMAIL"),
			Password: os.Getenv("PASSWORD"),
			Address:  os.Getenv("ADDRESS"),
		},
	}
}
