package httpserver

import (
	"context"
	"errors"
	"fmt"
	"net"
	"net/http"

	"go.uber.org/zap"
)

type Component struct {
	address string
	log     *zap.Logger
	server  *http.Server
}

func NewComponent(opts ...Option) *Component {
	options := defaultOptions()
	for _, opt := range opts {
		opt(options)
	}

	address := net.JoinHostPort(options.host, options.port)

	log := options.log.With(zap.String("component", "httpServer"))

	log.Info("initializing http server with configuration",
		zap.String("address", address),
		zap.Duration("readTimeout", options.readTimeout),
		zap.Duration("writeTimeout", options.writeTimeout),
		zap.Duration("idleTimeout", options.idleTimeout),
	)

	return &Component{
		address: address,
		server: &http.Server{
			Handler:      options.handler,
			ReadTimeout:  options.readTimeout,
			WriteTimeout: options.writeTimeout,
			IdleTimeout:  options.idleTimeout,
		},
		log: log,
	}
}

func (c *Component) Start(ctx context.Context) error {
	c.log.Info("starting http server")

	lisCfg := &net.ListenConfig{}
	lis, err := lisCfg.Listen(ctx, "tcp", c.address)
	if err != nil {
		return fmt.Errorf("cannot listen expected address %s: %w", c.address, err)
	}

	errCh := make(chan error)
	go func() {
		select {
		case errCh <- c.server.Serve(lis):
		case <-ctx.Done():
		}
		close(errCh)
	}()

	select {
	case sErr := <-errCh:
		if errors.Is(sErr, http.ErrServerClosed) {
			return nil
		}

		return fmt.Errorf("cannot serve expected address %s: %w", c.address, sErr)

	case <-ctx.Done():
		return nil
	}
}

func (c *Component) Stop(ctx context.Context) error {
	c.log.Info("stoping http server")

	return c.server.Shutdown(ctx)
}
