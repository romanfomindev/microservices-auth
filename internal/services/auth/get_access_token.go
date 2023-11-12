package auth

import (
	"context"

	"github.com/romanfomindev/microservices-auth/internal/utils"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *authService) GetAccessToken(ctx context.Context, refreshToken string) (string, error) {
	claims, err := utils.VerifyToken(refreshToken, []byte(s.refreshTokenSecretKey))
	if err != nil {
		return "", status.Errorf(codes.Unauthenticated, "invalid token: %s", err.Error())
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
