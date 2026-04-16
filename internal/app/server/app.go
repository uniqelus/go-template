package serverapp

import (
	"context"
	"fmt"

	"go.uber.org/zap"
	"golang.org/x/sync/errgroup"

	serverhttp "go-template/internal/transport/server/http"
	httpserver "go-template/pkg/components/http/server"
)

type App struct {
	log        *zap.Logger
	httpServer *httpserver.Component
}

func NewApp(cfg *Config, log *zap.Logger) *App {
	log = log.With(zap.String("service", "server"))

	return &App{
		log: log,
		httpServer: httpserver.NewComponent(
			httpserver.WithHost(cfg.HttpServer.Host),
			httpserver.WithPort(cfg.HttpServer.Port),
			httpserver.WithHandler(serverhttp.NewRouter(log)),
			httpserver.WithReadTimeout(cfg.HttpServer.ReadTimeout),
			httpserver.WithWriteTimeout(cfg.HttpServer.WriteTimeout),
			httpserver.WithIdleTimeout(cfg.HttpServer.IdleTimeout),
			httpserver.WithLogger(log),
		),
	}
}

func (a *App) Start(ctx context.Context) error {
	g, gCtx := errgroup.WithContext(ctx)

	g.Go(func() error {
		return a.httpServer.Start(gCtx)
	})

	if err := g.Wait(); err != nil {
		a.log.Error("failed to start appliction", zap.Error(err))
		return fmt.Errorf("fatal error during application start: %w", err)
	}

	return nil
}

func (a *App) Stop(ctx context.Context) error {
	g, gCtx := errgroup.WithContext(ctx)

	g.Go(func() error {
		return a.httpServer.Stop(gCtx)
	})

	if err := g.Wait(); err != nil {
		a.log.Error("failed to stop appliction", zap.Error(err))
		return fmt.Errorf("fatal error during application stop: %w", err)
	}

	return nil
}
