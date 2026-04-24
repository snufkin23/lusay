package service

import (
	"fmt"

	"github.com/snufkin23/lucisay/internal/core/ports"
)

// AIService orchestrates AI related use cases
type AIService struct {
	provider ports.AIProvider
}

// NewAIService creates a new instance of AIService
func NewAIService(provider ports.AIProvider) *AIService {
	return &AIService{
		provider: provider,
	}
}

// GenerateResponse takes a user prompt and returns a response from the AI provider
func (s *AIService) GenerateResponse(prompt string) (string, error) {
	if prompt == "" {
		return "", fmt.Errorf("service.GenerateResponse: prompt cannot be empty")
	}

	resp, err := s.provider.Generate(prompt)
	if err != nil {
		return "", fmt.Errorf("service.GenerateResponse: %w", err)
	}
	return resp.Content, nil
}
