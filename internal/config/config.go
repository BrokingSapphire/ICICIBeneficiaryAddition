package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
    Server ServerConfig
    ICICI  ICICIConfig
}

type ServerConfig struct {
    Port string
    Env  string
}

type ICICIConfig struct {
    BaseURL   string
    APIKey    string
    APISecret string
}

func Load() *Config {
    // Load .env file
    if err := godotenv.Load(); err != nil {
        log.Println("No .env file found, using system environment variables")
    }

    return &Config{
        Server: ServerConfig{
            Port: getEnv("PORT", "8080"),
            Env:  getEnv("GO_ENV", "development"),
        },
        ICICI: ICICIConfig{
            BaseURL:   getEnv("ICICI_BASE_URL", "https://api.icicidirect.com"),
            APIKey:    getEnv("ICICI_API_KEY", ""),
            APISecret: getEnv("ICICI_API_SECRET", ""),
        },
    }
}

func getEnv(key, defaultValue string) string {
    if value := os.Getenv(key); value != "" {
        return value
    }
    return defaultValue
}