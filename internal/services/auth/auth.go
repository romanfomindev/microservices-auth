package auth

import (
	"time"

	"github.com/romanfomindev/microservices-auth/internal/repositories"
	"github.com/romanfomindev/microservices-auth/internal/services"
)

type authService struct {
	refreshTokenSecretKey  string
	accessTokenSecretKey   string
	refreshTokenExpiration time.Duration
	accessTokenExpiration  time.Duration
	repo                   repositories.UserRepository
}

func NewAuthService(refreshTokenSecretKey, accessTokenSecretKey string, refreshTokenExpiration, accessTokenExpiration time.Duration, repo repositories.UserRepository) services.AuthService {
	return &authService{
		accessTokenSecretKey:   accessTokenSecretKey,
		refreshTokenSecretKey:  refreshTokenSecretKey,
		refreshTokenExpiration: refreshTokenExpiration,
		accessTokenExpiration:  accessTokenExpiration,
		repo:                   repo,
	}
}
