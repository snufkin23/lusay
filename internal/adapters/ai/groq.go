package ai

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/snufkin23/lusay/internal/adapters/config"
	"github.com/snufkin23/lusay/internal/core/domain"
	"github.com/snufkin23/lusay/internal/core/ports"
)

var _ ports.AIProvider = (*GroqClient)(nil)

// GroqClient is the adapter for Groq AI API
type GroqClient struct {
	cfg    *config.Config
	client *http.Client
}

func NewGroqClient(cfg *config.Config) *GroqClient {
	return &GroqClient{
		cfg:    cfg,
		client: &http.Client{},
	}
}

// Generate calls the Groq API to generate a response for the given prompt
func (c *GroqClient) Generate(prompt string) (*domain.AIResponse, error) {
	body := GroqRequest{
		Model: c.cfg.Model,
		Messages: []GroqMessage{
			{
				Role:    "user",
				Content: prompt,
			},
		},
	}

	jsonBody, err := json.Marshal(body)
	if err != nil {
		return nil, fmt.Errorf("ai.Generate marshal: %w", err)
	}

	req, err := http.NewRequest(
		"POST",
		c.cfg.GroqBaseURL,
		bytes.NewBuffer(jsonBody),
	)
	if err != nil {
		return nil, fmt.Errorf("ai.Generate request: %w", err)
	}

	req.Header.Set("Authorization", "Bearer "+c.cfg.GroqAPIKey)
	req.Header.Set("Content-Type", "application/json")

	res, err := c.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("ai.Generate do: %w", domain.ErrNetworkFailure)
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		respBody, _ := io.ReadAll(res.Body)
		switch res.StatusCode {
		case http.StatusTooManyRequests:
			return nil, fmt.Errorf("ai.Generate rate limit: %w", domain.ErrRateLimitExceeded)
		case http.StatusBadRequest:
			return nil, fmt.Errorf("ai.Generate bad request: %w", domain.ErrInvalidInput)
		case http.StatusForbidden:
			return nil, fmt.Errorf("ai.Generate forbidden: %w", domain.ErrContentFiltered)
		default:
			return nil, fmt.Errorf("ai.Generate non-200 (%d): %s", res.StatusCode, string(respBody))
		}
	}

	var groqResp GroqResponse
	if err := json.NewDecoder(res.Body).Decode(&groqResp); err != nil {
		return nil, fmt.Errorf("ai.Generate decode: %w", err)
	}

	if len(groqResp.Choices) == 0 {
		return nil, fmt.Errorf("ai.Generate no choices: %w", domain.ErrEmptyResponse)
	}

	content := groqResp.Choices[0].Message.Content
	if content == "" {
		return nil, fmt.Errorf("ai.Generate empty content: %w", domain.ErrEmptyResponse)
	}

	return &domain.AIResponse{
		Content: content,
	}, nil
}
