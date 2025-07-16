package env

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

// Env holds basic environment configuration
var Env *Environment

type Environment struct {
	Port    string
	Env     string
	APIPath string
}

// Load initializes the environment configuration
func Load() {
	// Load .env file
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using system environment variables")
	}

	// Initialize the global Env variable
	Env = &Environment{
		Port:    getEnv("PORT", "8080"),
		Env:     getEnv("GO_ENV", "development"),
		APIPath: getEnv("API_PATH", "/api/v1"),
	}
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
