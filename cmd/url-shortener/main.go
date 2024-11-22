package main

import (
	"log/slog"
	"os"
	"os/signal"
	"syscall"
	"url-shortener/internal/app"
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

	log := NewLogger(config.Env)

	app := app.New(log, config.HTTPServer) //TODO: add arguments for storage, cache and etc.

	go func() {
		app.HTTPServer.MustRun()
	}()

	// Graceful shutdown
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGTERM, syscall.SIGINT)

	<-stop

	app.HTTPServer.Run()
	log.Info("Gracefully stopped!")
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
