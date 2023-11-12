package auth_v1

import (
	"context"

	"github.com/romanfomindev/microservices-auth/internal/services"
	desc "github.com/romanfomindev/microservices-auth/pkg/auth_v1"
)

type AuthV1Handlers struct {
	desc.UnimplementedAuthServiceServer
	serv services.AuthService
}

func NewAuthHandlers(service services.AuthService) *AuthV1Handlers {
	return &AuthV1Handlers{
		serv: service,
	}
}

func (h *AuthV1Handlers) Login(ctx context.Context, request *desc.LoginRequest) (*desc.LoginResponse, error) {
	refreshToken, err := h.serv.Login(ctx, request.GetEmail(), request.GetPassword())
	if err != nil {
		return nil, err
	}

	return &desc.LoginResponse{RefreshToken: refreshToken}, err
}

func (h *AuthV1Handlers) GetRefreshToken(ctx context.Context, request *desc.GetRefreshTokenRequest) (*desc.GetRefreshTokenResponse, error) {
	newRefreshToken, err := h.serv.GetRefreshToken(ctx, request.OldRefreshToken)
	if err != nil {
		return nil, err
	}

	return &desc.GetRefreshTokenResponse{RefreshToken: newRefreshToken}, err

}

func (h *AuthV1Handlers) GetAccessToken(ctx context.Context, request *desc.GetAccessTokenRequest) (*desc.GetAccessTokenResponse, error) {
	accessToken, err := h.serv.GetAccessToken(ctx, request.RefreshToken)
	if err != nil {
		return nil, err
	}

	return &desc.GetAccessTokenResponse{AccessToken: accessToken}, err
}
