package ai

// GroqMessage defines a message in the Groq AI API.
type GroqMessage struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

// GroqRequest defines the request body for the Groq AI API.
// This is a DTO (Data Transfer Object) specific to the Groq adapter.
type GroqRequest struct {
	Model    string         `json:"model"`
	Messages []GroqMessage `json:"messages"`
}

// GroqChoice defines a single choice from the Groq AI API response.
type GroqChoice struct {
	Message GroqMessage `json:"message"`
}

// GroqResponse defines the response body from the Groq AI API.
// This is a DTO specific to the Groq adapter.
type GroqResponse struct {
	Choices []GroqChoice `json:"choices"`
}
