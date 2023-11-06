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
