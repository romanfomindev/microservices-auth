package auth

import (
	"context"

	"github.com/romanfomindev/microservices-auth/internal/utils"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *authService) GetRefreshToken(ctx context.Context, refreshToken string) (string, error) {
	claims, err := utils.VerifyToken(refreshToken, []byte(s.refreshTokenSecretKey))
	if err != nil {
		return "", status.Errorf(codes.Aborted, "invalid refresh token")
	}

	user, err := s.repo.GetByEmail(ctx, claims.Email)
	if err != nil {
		return "", err
	}

	newRefreshToken, err := utils.GenerateToken(*user, []byte(s.refreshTokenSecretKey), s.refreshTokenExpiration)
	if err != nil {
		return "", err
	}

	return newRefreshToken, nil
}
