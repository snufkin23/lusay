package main

import (
	"os"

	"github.com/snufkin23/lusay/internal/adapters/ai"
	"github.com/snufkin23/lusay/internal/adapters/cli"
	"github.com/snufkin23/lusay/internal/adapters/config"
	"github.com/snufkin23/lusay/internal/core/ports"
	"github.com/snufkin23/lusay/internal/core/service"
	"github.com/snufkin23/lusay/internal/utils/logger"
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
	var provider ports.AIProvider
	
	// Check if we should use the server client or the direct Groq client
	if os.Getenv("LUSAY_SERVER_URL") != "" {
		provider = ai.NewServerClient(os.Getenv("LUSAY_SERVER_URL"))
	} else {
		provider = ai.NewGroqClient(cfg)
	}

	// 1. Apply Guard Middleware (Filter harmful input)
	provider = service.NewGuardProvider(provider)

	// 2. Apply Persona Middleware (Inject cat personality)
	provider = service.NewPersonaProvider(provider)

	aiSvc := service.NewAIService(provider)

	// Initialize the CLI adapter
	app := cli.NewApp(aiSvc, l)

	// Start the application
	app.Run()
}
