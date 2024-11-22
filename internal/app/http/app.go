package httpapp

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"url-shortener/internal/config"
)

type App struct {
	log        *slog.Logger
	httpServer *http.Server
}

func New(
	log *slog.Logger,
	// Redirect and shortener services,
	srv config.HTTPServer,
) *App {
	httpServer := &http.Server{
		Addr: fmt.Sprintf("%s:%d", srv.Host, srv.Port),
		ReadTimeout: srv.ReadTimeout,
		IdleTimeout: srv.IdleTimeout,
	}

	// TODO: here i should register handlers, middlewares and etc.

	return &App{
		log: log,
		httpServer: httpServer,
		// smth else,
	}
}

func (a *App) MustRun() {
	if err := a.Run(); err != nil {
		panic(err)
	}
}

func (a *App) Run() error {
	const op = "httpapp.Run"

	a.log.With(
		slog.String("op", op),
	)

	a.log.Info(
		"HTTP server started",
		slog.String("addr", a.httpServer.Addr),
	)

	if err := a.httpServer.ListenAndServe(); err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}

func (a *App) Stop() {
	const op = "httpapp.Stop"

	a.log.With(slog.String("op", op)).
		Info(
			"stopping HTTP server",
			slog.String("addr", a.httpServer.Addr),
		)

	// TODO: this function returns error, idk what to do with it. To find out...
	a.httpServer.Shutdown(context.Background())
}
