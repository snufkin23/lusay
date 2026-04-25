package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/snufkin23/lusay/internal/adapters/ai"
	"github.com/snufkin23/lusay/internal/adapters/config"
	httpAdapter "github.com/snufkin23/lusay/internal/adapters/http"
	"github.com/snufkin23/lusay/internal/core/service"
	"github.com/snufkin23/lusay/internal/utils/logger"
)

func main() {
	l := logger.New()

	cfg, err := config.LoadConfig()
	if err != nil {
		l.Fatal("failed to load config", err)
	}

	// Core wiring
	groqClient := ai.NewGroqClient(cfg)
	aiSvc := service.NewAIService(groqClient)

	// HTTP Adapter
	handler := httpAdapter.NewHandler(aiSvc, l.Slog()) // Assuming logger.New() returns a wrapper with Slog()
	
	server := &http.Server{
		Addr:    ":8080",
		Handler: handler.Router(),
	}

	// Graceful shutdown
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

	go func() {
		l.Raw("🚀 lusay server starting on :8080")
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			l.Fatal("server failed", err)
		}
	}()

	<-stop

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		l.Fatal("server forced to shutdown", err)
	}

	l.Raw("👋 server stopped")
}
