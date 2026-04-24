package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/briandowns/spinner"
	"github.com/snufkin23/lucisay/internal/adapters/ai"
	"github.com/snufkin23/lucisay/internal/adapters/config"
	"github.com/snufkin23/lucisay/internal/core/service"
	"github.com/snufkin23/lucisay/internal/utils/catsay"
	"github.com/snufkin23/lucisay/internal/utils/logger"
)

func main() {
	// initialize logger
	l := logger.New()

	// load environment configuration
	cfg, err := config.LoadConfig()
	if err != nil {
		l.Fatal("failed to load config", err)
	}

	// Composition Root: wiring adapters to services
	groqClient := ai.NewGroqClient(cfg)
	aiSvc := service.NewAIService(groqClient)

	// Configure the loading animation (spinner)
	s := spinner.New(spinner.CharSets[11], 100*time.Millisecond)
	s.Suffix = " 🐱 lucisay is thinking..."
	s.Color = "\033[36m" // Cyan color

	// Interactive mode: REPL (Read-Eval-Print Loop)
	fmt.Println("🐱 Welcome to lucisay! (Type 'exit' or 'quit' to leave)")
	
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("\nYou: ")
		if !scanner.Scan() {
			break
		}

		userInput := strings.TrimSpace(scanner.Text())
		if userInput == "" {
			continue
		}

		// Check for exit commands
		lowerInput := strings.ToLower(userInput)
		if lowerInput == "exit" || lowerInput == "quit" {
			fmt.Println("🐱 Bye bye!")
			break
		}

		// Start loading animation
		s.Start()

		// Execute use case: get AI response
		resp, err := aiSvc.GenerateResponse(userInput)
		
		// Stop loading animation immediately after response
		s.Stop()

		if err != nil {
			l.Error("failed to generate response", err)
			continue
		}

		// Format the response using the cat-say formatter
		formattedOutput := catsay.Format(resp)
		l.Raw(formattedOutput)
	}

	if err := scanner.Err(); err != nil {
		l.Fatal("error reading input", err)
	}
}
