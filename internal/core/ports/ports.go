package ports

import "github.com/snufkin23/lusay/internal/core/domain"

// AIProvider is the port for interacting with an AI service
type AIProvider interface {
	// Generate takes a prompt and returns an AI response
	Generate(prompt string) (*domain.AIResponse, error)
}
