package main

import (
	"fmt"

	"github.com/snufkin23/lucisay/adapter/domain/config"
	"github.com/snufkin23/lucisay/adapter/utils/logger"
)

func main() {
	l := logger.New()
	l.Info("Starting lucisay...")

	cfg := config.LoadConfig()
	if cfg.GroqAPIKey == "" {
		l.Error("GROQ_API_KEY is not set", nil)
	} else {
		l.Info("Config loaded successfully")
	}

	fmt.Println("Hello from lucisay!")
}
