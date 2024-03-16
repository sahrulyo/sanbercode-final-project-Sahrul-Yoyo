package config

import (
	"fmt"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

// Config holds configuration variables
type Config struct {
	DBHost     string `envconfig:"DB_HOST"`
	DBUser     string `envconfig:"DB_USER"`
	DBPassword string `envconfig:"DB_PASSWORD"`
	DBName     string `envconfig:"DB_NAME"`
	ServerPort string `envconfig:"SERVER_PORT"`
	JWTSecret  string `envconfig:"JWT_SECRET"`
}

// LoadConfig loads configuration from environment variables or a .env file
func LoadConfig() (*Config, error) {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file:", err)
	}

	cfg := &Config{}
	err = envconfig.Process("", cfg)
	if err != nil {
		return nil, err
	}

	return cfg, nil
}
