package env

import (
	"errors"
	"os"
	"time"

	"github.com/romanfomindev/microservices-auth/internal/config"
)

const (
	refreshTokenSecretKey  = "REFRESH_TOKEN_SECRET_KEY"
	accessTokenSecretKey   = "ACCESS_TOKEN_SECRET_KEY"
	refreshTokenExpiration = "REFRESH_TOKEN_EXPIRATION"
	accessTokenExpiration  = "ACCESS_TOKEN_EXPIRATION"
)

type authConfig struct {
	refreshTokenSecretKey  string
	accessTokenSecretKey   string
	refreshTokenExpiration time.Duration
	accessTokenExpiration  time.Duration
}

func NewAuthConfig() (config.AuthConfig, error) {
	refreshToken := os.Getenv(refreshTokenSecretKey)
	if len(refreshToken) == 0 {
		return nil, errors.New("refresh token not found")
	}
	accessToken := os.Getenv(accessTokenSecretKey)
	if len(refreshToken) == 0 {
		return nil, errors.New("access token not found")
	}
	refreshTokenExpirationEnv := os.Getenv(refreshTokenExpiration)
	if len(refreshTokenExpirationEnv) == 0 {
		return nil, errors.New("refresh token expiration not found")
	}
	refreshTokenExpirationValue, err := time.ParseDuration(refreshTokenExpirationEnv)
	if err != nil {
		return nil, errors.New("invalid refresh token expiration")
	}
	accessTokenExpirationEnv := os.Getenv(accessTokenExpiration)
	if len(accessTokenExpirationEnv) == 0 {
		return nil, errors.New("access token expiration not found")
	}
	accessTokenExpirationValue, err := time.ParseDuration(accessTokenExpirationEnv)
	if err != nil {
		return nil, errors.New("invalid refresh token expiration")
	}

	return &authConfig{
		refreshTokenSecretKey:  refreshToken,
		accessTokenSecretKey:   accessToken,
		refreshTokenExpiration: refreshTokenExpirationValue,
		accessTokenExpiration:  accessTokenExpirationValue,
	}, nil
}

func (cfg *authConfig) RefreshTokenSecretKey() string {
	return cfg.refreshTokenSecretKey
}

func (cfg *authConfig) AccessTokenSecretKey() string {
	return cfg.accessTokenSecretKey
}

func (cfg *authConfig) RefreshTokenExpiration() time.Duration {
	return cfg.refreshTokenExpiration
}

func (cfg *authConfig) AccessTokenExpiration() time.Duration {
	return cfg.accessTokenExpiration
}
