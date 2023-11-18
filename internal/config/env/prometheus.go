package env

import (
	"errors"
	"net"
	"os"

	"github.com/romanfomindev/microservices-auth/internal/config"
)

const (
	HOST = "PROMETHEUS_HOST"
	PORT = "PROMETHEUS_PORT"
)

type prometheusConfig struct {
	host string
	port string
}

func NewPrometheusConfig() (config.PrometheusConfig, error) {
	host := os.Getenv(HOST)
	if len(host) == 0 {
		return nil, errors.New("prometheus host not found")
	}

	port := os.Getenv(PORT)
	if len(port) == 0 {
		return nil, errors.New("prometheus port not found")
	}

	return prometheusConfig{
		host: host,
		port: port,
	}, nil
}

func (cfg prometheusConfig) Address() string {
	return net.JoinHostPort(cfg.host, cfg.port)
}
