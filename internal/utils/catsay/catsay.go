package catsay

import (
	"strings"
)

// CatResponse holds the components of a cat's response
type CatResponse struct {
	Stages []Stage
	Art    string
	Mood   string
}

type Stage struct {
	Header  string
	Content string
}

// Format parses the persona-generated text into structured stages
func Format(text string, mood string) CatResponse {
	stages := []Stage{}
	lines := strings.Split(text, "\n")

	var currentHeader string
	var currentContent strings.Builder

	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}
		if strings.Contains(line, "💭 THE DAYDREAM:") {
			if currentHeader != "" {
				stages = append(stages, Stage{Header: currentHeader, Content: strings.TrimSpace(currentContent.String())})
				currentContent.Reset()
			}
			currentHeader = "💭 THE DAYDREAM"
			rest := strings.TrimSpace(strings.TrimPrefix(line, "💭 THE DAYDREAM:"))
			currentContent.WriteString(rest)
		} else if strings.Contains(line, "🐾 THE MEOW:") {
			if currentHeader != "" {
				stages = append(stages, Stage{Header: currentHeader, Content: strings.TrimSpace(currentContent.String())})
				currentContent.Reset()
			}
			currentHeader = "🐾 THE MEOW"
			rest := strings.TrimSpace(strings.TrimPrefix(line, "🐾 THE MEOW:"))
			currentContent.WriteString(rest)
		} else if strings.Contains(line, "💤 THE NAP:") {
			if currentHeader != "" {
				stages = append(stages, Stage{Header: currentHeader, Content: strings.TrimSpace(currentContent.String())})
				currentContent.Reset()
			}
			currentHeader = "💤 THE NAP"
			rest := strings.TrimSpace(strings.TrimPrefix(line, "💤 THE NAP:"))
			currentContent.WriteString(rest)
		} else {
			if currentContent.Len() > 0 {
				currentContent.WriteString(" ")
			}
			currentContent.WriteString(line)
		}
	}

	if currentHeader != "" {
		stages = append(stages, Stage{Header: currentHeader, Content: strings.TrimSpace(currentContent.String())})
	}

	return CatResponse{
		Stages: stages,
		Art:    getDetailedCat(mood),
		Mood:   mood,
	}
}

func getDetailedCat(mood string) string {
	switch strings.ToUpper(mood) {
	case "HAPPY":
		// Blissful, half-lidded, tail curled high
		return `   /\_____/\
  /  ^   ^  \
 ( (  uwu  ) )
  \ ` + "`" + `~~~~~` + "`" + ` /
  /         \
 |   |   |   |
  \___|___|__/
   (🐾) (🐾)`

	case "SHOCKED":
		// Wide-eyed, spiky fur, raised paws mid-air
		return `  ^ /\_/|^ ^
 /| O   O |\
| |  !!!  | |
 \|  ___  |/
  / /   \ \
 / / \_/ \ \
|_|       |_|
 (🐾)   (🐾)`

	case "NERD":
		// Glasses on, one paw raised like a professor
		return `   /\_____/\
  / -o---o- \
 |  (=====)  |
 |   \___/   |
  \  [===]  /
  /  |   |  \
 |   |   |   |
  \___|___|__/
   (🐾) (🐾)`

	case "SNEAKY":
		// Crouched low, one eye barely open, tail coiled
		return `    ___
  /     \__
 | -   .   )
  \  ~~~  /
  / ___  /~~~~~)
 |_|   |_|
  (🐾)(🐾)`

	case "HISSING":
		// Arched back, spiky spine, bared fangs
		return `  ^ /\_/\ ^
 ^| >   < |^
 ^|  VVV  |^
  \ ~~~~~ /
  /^ ^^^ ^\
 / ^^^^^^^  \
|_|         |_|
 (🐾)     (🐾)`

	default:
		return `   /\_____/\
  /  o   o  \
 (    ---    )
  \  ~~~~~  /
  /         \
 |   |   |   |
  \___|___|__/
   (🐾) (🐾)`
	}
}
