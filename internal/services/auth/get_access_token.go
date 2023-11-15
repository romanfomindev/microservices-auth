package auth

import (
	"context"
	"fmt"

	"github.com/romanfomindev/microservices-auth/internal/utils"
)

func (s *authService) GetAccessToken(ctx context.Context, refreshToken string) (string, error) {
	claims, err := utils.VerifyToken(refreshToken, []byte(s.refreshTokenSecretKey))
	if err != nil {
		return "", fmt.Errorf("invalid refresh token: %w", err)
	}

	user, err := s.repo.GetByEmail(ctx, claims.Email)
	if err != nil {
		return "", err
	}

	accessToken, err := utils.GenerateToken(*user, []byte(s.accessTokenSecretKey), s.accessTokenExpiration)
	if err != nil {
		return "", err
	}

	return accessToken, nil
}
