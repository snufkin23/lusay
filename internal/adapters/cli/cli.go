package cli

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/briandowns/spinner"
	"github.com/snufkin23/lucisay/internal/core/service"
	"github.com/snufkin23/lucisay/internal/utils/catsay"
	"github.com/snufkin23/lucisay/internal/utils/logger"
)

// App handles the CLI interaction for lucisay
type App struct {
	aiSvc  *service.AIService
	logger *logger.Logger
}

// NewApp creates a new instance of the CLI application
func NewApp(aiSvc *service.AIService, logger *logger.Logger) *App {
	return &App{
		aiSvc:  aiSvc,
		logger: logger,
	}
}

// Run starts the interactive REPL loop
func (a *App) Run() {
	// Configure the loading animation (spinner)
	s := spinner.New(spinner.CharSets[11], 100*time.Millisecond)
	s.Suffix = " 🐱 lucisay is thinking..."

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
		resp, err := a.aiSvc.GenerateResponse(userInput)

		// Stop loading animation immediately after response
		s.Stop()

		if err != nil {
			a.logger.Error("failed to generate response", err)
			continue
		}

		// Format the response using the cat-say formatter
		formattedOutput := catsay.Format(resp)
		a.logger.Raw(formattedOutput)
	}

	if err := scanner.Err(); err != nil {
		a.logger.Fatal("error reading input", err)
	}
}
