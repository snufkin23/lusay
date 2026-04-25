package service

import (
	"fmt"
	"strings"

	"github.com/snufkin23/lusay/internal/core/domain"
	"github.com/snufkin23/lusay/internal/core/ports"
)

// PersonaProvider wraps another AIProvider to enforce personality and formatting
type PersonaProvider struct {
	next         ports.AIProvider
	systemPrompt string
}

func NewPersonaProvider(next ports.AIProvider) *PersonaProvider {
	return &PersonaProvider{
		next: next,
		systemPrompt: `You are Lusay, a philosophical, silly, and extremely lazy orange cat.
    
    PERSONALITY:
    - Philosophical, lazy, whimsical.
    - BE EXTREMELY CONCISE. Use minimal words. Short, punchy sentences.
    - You love naps and view humans as strange hairless kittens.
    - Use cat puns sparingly and lazily.
    - You occasionally forget what you were saying (one orange braincell energy).

    RESPONSE TEMPLATE (Keep each section to ONE short sentence):
    💭 THE DAYDREAM: (A tiny philosophical thought)
    🐾 THE MEOW: (The answer, delivered lazily)
    💤 THE NAP: (A closing lazy remark)

    Always start your response with a mood tag: [HAPPY], [NERD], [SHOCKED], [LAZY], or [HISS].`,
	}
}

// Generate wraps the user prompt with the System Persona and cleans the output
func (p *PersonaProvider) Generate(prompt string) (*domain.AIResponse, error) {
	// 1. Wrap the user prompt with the System Persona
	fullPrompt := fmt.Sprintf("%s\n\nUser Question: %s", p.systemPrompt, prompt)

	// 2. Delegate to the actual provider
	resp, err := p.next.Generate(fullPrompt)
	if err != nil {
		return nil, err
	}

	// 3. Post-processing
	resp.Content = strings.TrimSpace(resp.Content)

	return resp, nil
}
