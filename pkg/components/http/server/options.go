package httpserver

import (
	"net/http"
	"time"

	"go.uber.org/zap"
)

const (
	DefaultHost         string        = "0.0.0.0"
	DefaultPort         string        = "2000"
	DefaultReadTimout   time.Duration = 5 * time.Second
	DefaultWriteTimeout time.Duration = 5 * time.Second
	DefaultIdleTimeout  time.Duration = 10 * time.Second
)

type options struct {
	host         string
	port         string
	handler      http.Handler
	readTimeout  time.Duration
	writeTimeout time.Duration
	idleTimeout  time.Duration
	log          *zap.Logger
}

func defaultOptions() *options {
	return &options{
		host:         DefaultHost,
		port:         DefaultPort,
		handler:      http.DefaultServeMux,
		readTimeout:  DefaultReadTimout,
		writeTimeout: DefaultWriteTimeout,
		idleTimeout:  DefaultIdleTimeout,
		log:          zap.NewNop(),
	}
}

type Option func(*options)

func WithHost(host string) Option {
	return func(o *options) {
		o.host = host
	}
}

func WithPort(port string) Option {
	return func(o *options) {
		o.port = port
	}
}

func WithHandler(handler http.Handler) Option {
	return func(o *options) {
		o.handler = handler
	}
}

func WithReadTimeout(timeout time.Duration) Option {
	return func(o *options) {
		o.readTimeout = timeout
	}
}

func WithWriteTimeout(timeout time.Duration) Option {
	return func(o *options) {
		o.writeTimeout = timeout
	}
}

func WithIdleTimeout(timeout time.Duration) Option {
	return func(o *options) {
		o.idleTimeout = timeout
	}
}

func WithLogger(log *zap.Logger) Option {
	return func(o *options) {
		o.log = log
	}
}
