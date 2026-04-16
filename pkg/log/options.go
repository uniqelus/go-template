package pkglog

type options struct {
	level       string
	outputPaths []string
}

func defaultOptions() *options {
	return &options{
		level:       "debug",
		outputPaths: []string{"stdout"},
	}
}

type Option func(*options)

func WithLevel(level string) Option {
	return func(o *options) {
		o.level = level
	}
}

func WithOutputPaths(paths ...string) Option {
	return func(o *options) {
		o.outputPaths = paths
	}
}
