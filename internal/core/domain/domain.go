package domain

import "errors"

// sentinel errors for business logic
var (
	ErrAIProviderFailure = errors.New("ai provider failed to generate response")
	ErrInvalidInput      = errors.New("invalid input provided")
)

// AIResponse holds the result from the AI provider
type AIResponse struct {
	Content string
}
