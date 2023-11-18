package env

import (
	"errors"
	"os"

	"github.com/romanfomindev/microservices-auth/internal/config"
)

const LEVEL = "LOG_LEVEL"

type loggerConfig struct {
	level string
}

func NewLoggerConfig() (config.LoggerConfig, error) {
	level := os.Getenv(LEVEL)
	if len(level) == 0 {
		return nil, errors.New("logger level env not found")
	}
	return &loggerConfig{
		level: level,
	}, nil
}

func (cfg *loggerConfig) Level() string {
	return cfg.level
}
