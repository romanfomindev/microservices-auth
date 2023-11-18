package config

import (
	"time"

	"github.com/joho/godotenv"
)

func Load(path string) error {
	err := godotenv.Load(path)
	if err != nil {
		return err
	}

	return nil
}

type GRPCConfig interface {
	Address() string
}

type PGConfig interface {
	DSN() string
	Timeout() time.Duration
}

type HTTPConfig interface {
	Address() string
}

type SwaggerConfig interface {
	Address() string
}

type AuthConfig interface {
	RefreshTokenSecretKey() string
	AccessTokenSecretKey() string
	RefreshTokenExpiration() time.Duration
	AccessTokenExpiration() time.Duration
}

type AppConfig interface {
	Env() string
}

type LoggerConfig interface {
	Level() string
}

type PrometheusConfig interface {
	Address() string
}
