package env

import (
	"errors"
	"os"

	"github.com/romanfomindev/microservices-auth/internal/config"
)

const ENV = "APP_ENV"

type appConfig struct {
	appEnv string
}

func NewAppConfig() (config.AppConfig, error) {
	appEnv := os.Getenv(ENV)
	if len(appEnv) == 0 {
		return nil, errors.New("app env not found")
	}

	return &appConfig{
		appEnv: appEnv,
	}, nil
}

func (cfg *appConfig) Env() string {
	return cfg.appEnv
}
