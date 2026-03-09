package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DB        string
	AppPort   string
	JWTSecret string
}

func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}

func Load() *Config {
	err := godotenv.Load()

	if err != nil {
		log.Println("Warning: .env file not found, using environment variables")
	}

	return &Config{
		DB:        getEnv("DB", "data.db"),
		AppPort:   getEnv("PORT", "8080"),
		JWTSecret: getEnv("JWT_SECRET", "mysecret"),
	}
}
