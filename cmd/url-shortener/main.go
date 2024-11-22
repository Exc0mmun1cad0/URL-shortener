package main

import (
	"fmt"
	"log/slog"
	"os"
	"url-shortener/internal/config"
	"url-shortener/pkg/badaslog"
)

const (
	envLocal = "local"
	envDev   = "dev"
	envProd  = "prod"
)

func main() {
	config := config.MustLoad()

	logger := NewLogger(config.Env)

	logger.Info(
		"starting url-shortener service",
		slog.String("env", config.Env),
		slog.String("address", fmt.Sprintf("%s:%d", config.HTTPServer.Host, config.HTTPServer.Port)),
	)

	// TODO: logger initialization

	// TODO: storage initialization

	// TODO: cache initialization

	// TODO: creating router
	// TODO: registering handlers
	// TODO: adding middlewares to router

	// TODO: HTTP server initialization
	// TODO: graceful shutdownm
}

func NewLogger(env string) *slog.Logger {
	var logger *slog.Logger

	switch env {
	case envLocal:
		logger = setupBadaSlogger()
	case envDev:
		logger = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
		)
	case envProd:
		logger = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}),
		)
	default:
		logger = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}),
		)
	}

	return logger
}

func setupBadaSlogger() *slog.Logger {
	var logger *slog.Logger

	opts := &slog.HandlerOptions{Level: slog.LevelDebug}
	logger = slog.New(badaslog.NewHandler(opts))

	return logger
}
