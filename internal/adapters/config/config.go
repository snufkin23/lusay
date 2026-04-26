package config

import (
	"errors"
	"os"

	"github.com/joho/godotenv"
)

const (
	envGroqAPIKey  string = "GROQ_API_KEY"
	envModel       string = "GROQ_MODEL"
	envGroqBaseURL string = "GROQ_BASE_URL"
)

type Config struct {
	GroqAPIKey  string
	Model       string
	GroqBaseURL string
}

func LoadConfig() (*Config, error) {
	_ = godotenv.Load()

	key := os.Getenv(envGroqAPIKey)
	if key == "" {
		return nil, errors.New("missing required env: GROQ_API_KEY")
	}

	model := os.Getenv(envModel)

	baseURL := os.Getenv(envGroqBaseURL)

	return &Config{
		GroqAPIKey:  key,
		Model:       model,
		GroqBaseURL: baseURL,
	}, nil
}
