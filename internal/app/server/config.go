package serverapp

type Config struct {
	Log LogConfig `yaml:"log"`
}

type LogConfig struct {
	Level       string   `yaml:"level" env:"LOG_LEVEL" env-default:"debug"`
	OutputPaths []string `yaml:"outputPaths" env:"LOG_OUTPUT_PATHS" env-default:"stdout,logs/app.log"`
}
