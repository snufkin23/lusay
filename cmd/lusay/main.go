package main

import (
	"github.com/snufkin23/lusay/internal/adapters/ai"
	"github.com/snufkin23/lusay/internal/adapters/cli"
	"github.com/snufkin23/lusay/internal/adapters/config"
	"github.com/snufkin23/lusay/internal/core/ports"
	"github.com/snufkin23/lusay/internal/core/service"
	"github.com/snufkin23/lusay/internal/utils/logger"
)

func main() {
	l := logger.New()

	cfg, err := config.LoadConfig()
	if err != nil {
		l.Fatal("failed to load config", err)
	}

	var provider ports.AIProvider = ai.NewGroqClient(cfg)

	provider = service.NewGuardProvider(provider)

	provider = service.NewPersonaProvider(provider)

	aiSvc := service.NewAIService(provider)

	app := cli.NewApp(aiSvc, l)

	app.Run()
}
