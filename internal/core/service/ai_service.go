package service

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/snufkin23/lusay/internal/core/domain"
	"github.com/snufkin23/lusay/internal/core/ports"
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

// GenerateResponse handles the flow from user prompt to mood-aware response
func (s *AIService) GenerateResponse(prompt string) (*domain.CatResponse, error) {
	if prompt == "" {
		return nil, fmt.Errorf("service.GenerateResponse: %w", domain.ErrInvalidInput)
	}

	resp, err := s.provider.Generate(prompt)
	if err != nil {
		return nil, fmt.Errorf("service.GenerateResponse: %w", err)
	}

	// Parse mood tag [MOOD] and clean text
	mood, cleanText := s.parseMood(resp.Content)

	return &domain.CatResponse{
		Text: cleanText,
		Mood: mood,
	}, nil
}

// parseMood extracts mood tag and returns clean content
func (s *AIService) parseMood(content string) (string, string) {
	re := regexp.MustCompile(`\[(HAPPY|NERD|SHOCKED|LAZY|HISS)\]`)
	matches := re.FindStringSubmatch(content)
	if len(matches) > 1 {
		return matches[1], strings.TrimSpace(re.ReplaceAllString(content, ""))
	}
	return "NEUTRAL", content
}
