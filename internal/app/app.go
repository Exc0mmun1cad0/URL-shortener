package app

import (
	"log/slog"
	httpapp "url-shortener/internal/app/http"
	"url-shortener/internal/config"
)

type App struct {
	HTTPServer *httpapp.App
}

func New(
	log *slog.Logger,
	server config.HTTPServer,
	// TODO: also data for postgres and redis connections
) *App {
	// TODO: db and cache initialization

	// TODO: Creating redirect and shortener api services.
	// TODO: Pass them db and cache conns and logger as arguments

	// TODO: initialize httpapp.App with those services and logger as arguments

	httpApp := httpapp.New(log, server)

	return &App{
		HTTPServer: httpApp,
	}
}
