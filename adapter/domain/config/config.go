package config

import (
	"os"

	"github.com/joho/godotenv"
)

const groqApiKeyEnv string = "GROQ_API_KEY"

type Config struct {
	GroqAPIKey string
}

func LoadConfig() *Config {
	_ = godotenv.Load()
	return &Config{
		GroqAPIKey: os.Getenv(groqApiKeyEnv),
	}
}
