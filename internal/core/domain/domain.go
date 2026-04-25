package domain

import "errors"

// sentinel errors for business logic
var (
	ErrAIProviderFailure = errors.New("ai provider failed to generate response")
	ErrInvalidInput      = errors.New("invalid input provided")
	ErrHarmfulContent    = errors.New("harmful or dangerous content detected")
	ErrNetworkFailure    = errors.New("network failure occurred while contacting ai provider")
	ErrRateLimitExceeded = errors.New("ai provider rate limit exceeded")
	ErrEmptyResponse     = errors.New("ai provider returned an empty response")
	ErrContentFiltered   = errors.New("response was filtered by ai provider safety settings")
)

// AIResponse holds the result from the AI provider
type AIResponse struct {
	Content string
}

// CatResponse holds a processed response with mood for the UI
type CatResponse struct {
	Text string
	Mood string
}
