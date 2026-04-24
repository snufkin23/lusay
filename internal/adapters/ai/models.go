package ai

// groqRequest defines the request body for the Groq AI API.
// This is a DTO (Data Transfer Object) specific to the Groq adapter.
type groqRequest struct {
	Model    string             `json:"model"`
	Messages []map[string]string `json:"messages"`
}

// groqResponse defines the response body from the Groq AI API.
// This is a DTO specific to the Groq adapter.
type groqResponse struct {
	Choices []struct {
		Message struct {
			Content string `json:"content"`
		} `json:"message"`
	} `json:"choices"`
}
