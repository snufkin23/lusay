package ai

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/snufkin23/lusay/internal/core/domain"
)

// ServerClient is an implementation of AIProvider that calls the Lusay Server
type ServerClient struct {
	baseURL    string
	httpClient *http.Client
}

func NewServerClient(baseURL string) *ServerClient {
	return &ServerClient{
		baseURL: baseURL,
		httpClient: &http.Client{
			Timeout: 30 * time.Second,
		},
	}
}

func (c *ServerClient) Generate(prompt string) (*domain.AIResponse, error) {
	reqBody := map[string]string{
		"prompt": prompt,
	}
	jsonBody, _ := json.Marshal(reqBody)

	req, err := http.NewRequest("POST", c.baseURL+"/api/v1/generate", bytes.NewBuffer(jsonBody))
	if err != nil {
		return nil, fmt.Errorf("server_client.Generate request: %w", err)
	}
	req.Header.Set("Content-Type", "application/json")

	res, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("server_client.Generate do: %w", err)
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("server_client.Generate non-200: %d", res.StatusCode)
	}

	var respBody struct {
		Content string `json:"content"`
	}
	if err := json.NewDecoder(res.Body).Decode(&respBody); err != nil {
		return nil, fmt.Errorf("server_client.Generate decode: %w", err)
	}

	return &domain.AIResponse{
		Content: respBody.Content,
	}, nil
}
