package serverapp

import "time"

type Config struct {
	Log        LogConfig        `yaml:"log"`
	HttpServer HttpServerConfig `yaml:"http_server"`
}

type LogConfig struct {
	Level       string   `yaml:"level" env:"LOG_LEVEL" env-default:"debug"`
	OutputPaths []string `yaml:"outputPaths" env:"LOG_OUTPUT_PATHS" env-default:"stdout"`
}

type HttpServerConfig struct {
	Host         string        `yaml:"host" env:"HTTP_SERVER_HOST" env-default:"127.0.0.1"`
	Port         string        `yaml:"port" env:"HTTP_SERVER_PORT" env-default:"7654"`
	ReadTimeout  time.Duration `yaml:"readTimeout" env:"HTTP_SERVER_READ_TIMEOUT" env-default:"3s"`
	WriteTimeout time.Duration `yaml:"writeTimeout" env:"HTTP_SERVER_WRITE_TIMEOUT" env-default:"3s"`
	IdleTimeout  time.Duration `yaml:"idleTimeout" env:"HTTP_SERVER_IDLE_TIMEOUT" env-default:"12s"`
}
