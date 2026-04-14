package pkglog

import (
	"fmt"

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
