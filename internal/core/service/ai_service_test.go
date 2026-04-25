package service

import (
	"errors"
	"testing"

	"github.com/snufkin23/lusay/internal/core/domain"
)

// MockAIProvider is a mock implementation of the AIProvider port
type MockAIProvider struct {
	GenerateFunc func(prompt string) (*domain.AIResponse, error)
}

func (m *MockAIProvider) Generate(prompt string) (*domain.AIResponse, error) {
	return m.GenerateFunc(prompt)
}

func TestGenerateResponse(t *testing.T) {
	tests := []struct {
		name          string
		prompt        string
		mockBehavior  func(prompt string) (*domain.AIResponse, error)
		expectedResp  string
		expectedError bool
	}{
		{
			name:   "success",
			prompt: "Hello",
			mockBehavior: func(prompt string) (*domain.AIResponse, error) {
				return &domain.AIResponse{Content: "Hi there!"}, nil
			},
			expectedResp:  "Hi there!",
			expectedError: false,
		},
		{
			name:   "provider error",
			prompt: "Hello",
			mockBehavior: func(prompt string) (*domain.AIResponse, error) {
				return nil, errors.New("api failure")
			},
			expectedResp:  "",
			expectedError: true,
		},
		{
			name:          "empty prompt",
			prompt:        "",
			mockBehavior:  nil, // should not be called
			expectedResp:  "",
			expectedError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mock := &MockAIProvider{GenerateFunc: tt.mockBehavior}
			svc := NewAIService(mock)

			resp, err := svc.GenerateResponse(tt.prompt)

			if (err != nil) != tt.expectedError {
				t.Errorf("GenerateResponse() error = %v, expectedError %v", err, tt.expectedError)
				return
			}
			if resp != tt.expectedResp {
				t.Errorf("GenerateResponse() = %v, expected %v", resp, tt.expectedResp)
			}
		})
	}
}
