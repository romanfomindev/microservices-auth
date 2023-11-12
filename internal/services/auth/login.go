package auth

import (
	"context"
	"errors"

	"github.com/romanfomindev/microservices-auth/internal/utils"
)

func (s *authService) Login(ctx context.Context, email, password string) (string, error) {
	user, err := s.repo.GetByEmail(ctx, email)
	if err != nil {
		return "", err
	}
	if !utils.VerifyPassword(user.Password, password) {
		return "", errors.New("invalid password")
	}

	// get refresh token
	refreshToken, err := utils.GenerateToken(*user, []byte(s.refreshTokenSecretKey), s.refreshTokenExpiration)
	if err != nil {
		return "", err
	}

	return refreshToken, nil
}
