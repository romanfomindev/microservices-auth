package user

import (
	"context"
	"github.com/romanfomindev/microservices-auth/internal/convertor"
	"github.com/romanfomindev/microservices-auth/internal/models"
)

func (s *service) Create(ctx context.Context, userService models.User) (uint64, error) {
	userRepo := convertor.ToUserCreateFromUserService(userService)
	return s.repo.Create(ctx, userRepo)
}
