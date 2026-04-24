package catsay

import (
	"fmt"
	"strings"
)

// Format wraps the text in a speech bubble and adds a cat ASCII art below it
func Format(text string) string {
	if text == "" {
		text = "..."
	}

	// Basic bubble formatting
	lines := strings.Split(text, "\n")
	var bubble strings.Builder

	// Top border
	bubble.WriteString("  ")
	bubble.WriteString(strings.Repeat("_", len(lines[0])+2))
	bubble.WriteString("\n")

	// Content
	for _, line := range lines {
		bubble.WriteString(fmt.Sprintf(" < %s >\n", line))
	}

	// Bottom border
	bubble.WriteString("    ")
	bubble.WriteString(strings.Repeat("-", len(lines[0])+2))
	bubble.WriteString("\n")

	// Cute Cat Art
	cat := `   |\---/|
   | o_o |
    \_^_/
    /   \
   /     \
   |  _  |
   | ( ) |
   \  _  /
    \___/`

	return bubble.String() + cat
}
