package user

import (
	"context"
	"github.com/romanfomindev/microservices-auth/internal/convertor"
	"github.com/romanfomindev/microservices-auth/internal/models"
)

func (s *service) Update(ctx context.Context, id uint64, userService models.User) error {
	userRepo := convertor.ToUserUpdateFromUserService(userService)
	return s.repo.Update(ctx, id, userRepo)
}
