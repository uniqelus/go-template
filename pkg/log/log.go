package pkglog

import (
	"fmt"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5/middleware"
	"go.uber.org/zap"
)

func NewLogger(opts ...Option) (*zap.Logger, error) {
	options := defaultOptions()
	for _, opt := range opts {
		opt(options)
	}

	level, err := zap.ParseAtomicLevel(options.level)
	if err != nil {
		return nil, fmt.Errorf("cannot parse atomic level %s: %w", options.level, err)
	}

	logCfg := zap.Config{
		Encoding:      "json",
		Level:         level,
		OutputPaths:   options.outputPaths,
		EncoderConfig: zap.NewProductionEncoderConfig(),
	}

	return logCfg.Build()
}

type LogFormatter struct {
	log *zap.Logger
}

func NewLogFormatter(log *zap.Logger) *LogFormatter {
	return &LogFormatter{
		log: log,
	}
}

func (f *LogFormatter) NewLogEntry(r *http.Request) middleware.LogEntry {
	return &LogEntry{
		LogFormatter: f,
		request:      r,
	}
}

type LogEntry struct {
	*LogFormatter
	request *http.Request
}

func (l *LogEntry) Write(status, bytes int, _ http.Header, elapsed time.Duration, _ interface{}) {
	l.log.Info("request completed",
		zap.String("method", l.request.Method),
		zap.String("path", l.request.URL.Path),
		zap.Int("status", status),
		zap.Int("bytes", bytes),
		zap.Duration("elapsed", elapsed),
	)
}

func (l *LogEntry) Panic(v interface{}, stack []byte) {
	l.log.Error("panic recovered",
		zap.String("method", l.request.Method),
		zap.String("path", l.request.URL.Path),
		zap.String("remote_addr", l.request.RemoteAddr),
		zap.String("panic_value", fmt.Sprintf("%v", v)),
		zap.ByteString("stack_trace", stack),
	)
}
