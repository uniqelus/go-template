package serverapp

import (
	"context"

	"go.uber.org/zap"
)

type App struct {
	log *zap.Logger
}

func NewApp(cfg *Config, log *zap.Logger) (*App, error) {
	log = log.With(zap.String("service", "server"))

	log.Debug("initializing application")
	log.Error("test error")

	return &App{
		log: log,
	}, nil
}

func Run(ctx context.Context) error {
	return nil
}

func start(ctx context.Context) error {
	return nil
}

func stop(ctx context.Context) error {
	return nil
}
