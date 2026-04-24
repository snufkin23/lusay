package main

import (
	"github.com/snufkin23/lucisay/internal/adapters/ai"
	"github.com/snufkin23/lucisay/internal/adapters/cli"
	"github.com/snufkin23/lucisay/internal/adapters/config"
	"github.com/snufkin23/lucisay/internal/core/service"
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

	// Initialize the CLI adapter
	app := cli.NewApp(aiSvc, l)

	// Start the application
	app.Run()
}
