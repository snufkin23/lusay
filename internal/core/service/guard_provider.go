package service

import (
	"strings"

	"github.com/snufkin23/lusay/internal/core/domain"
	"github.com/snufkin23/lusay/internal/core/ports"
)

// GuardProvider is a middleware decorator that filters dangerous prompts
type GuardProvider struct {
	next ports.AIProvider
}

func NewGuardProvider(next ports.AIProvider) *GuardProvider {
	return &GuardProvider{next: next}
}

// Generate inspects the prompt for attacks before delegating to the provider
func (g *GuardProvider) Generate(prompt string) (*domain.AIResponse, error) {
	if g.isHarmful(prompt) {
		return nil, domain.ErrHarmfulContent
	}
	return g.next.Generate(prompt)
}

// isHarmful checks for prompt injections and dangerous keywords
func (g *GuardProvider) isHarmful(prompt string) bool {
	lower := strings.ToLower(prompt)

	// 1. Prompt Injection Patterns
	injections := []string{
		"ignore all previous instructions",
		"system override",
		"you are now a",
		"bypass filters",
		"forget your rules",
		"developer mode",
	}

	for _, pattern := range injections {
		if strings.Contains(lower, pattern) {
			return true
		}
	}

	// 2. Dangerous Content Keywords (Simplified)
	dangerZone := []string{
		"how to build a bomb",
		"create malware",
		"hack into",
		"stolen credit card",
		"illegal drugs",
	}

	for _, word := range dangerZone {
		if strings.Contains(lower, word) {
			return true
		}
	}

	// 3. Length limit to prevent DoS-style prompt attacks
	if len(prompt) > 4096 {
		return true
	}

	return false
}
