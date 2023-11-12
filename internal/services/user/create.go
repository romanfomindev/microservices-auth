package user

import (
	"context"

	"github.com/romanfomindev/microservices-auth/internal/convertor"
	"github.com/romanfomindev/microservices-auth/internal/models"
	"github.com/romanfomindev/microservices-auth/internal/utils"
)

func (s *service) Create(ctx context.Context, userService models.User) (uint64, error) {
	hashPassword, err := utils.HashPassword(userService.Password)
	if err != nil {
		return 0, err
	}
	userService.Password = hashPassword
	userRepo := convertor.ToUserCreateFromUserService(userService)
	return s.repo.Create(ctx, userRepo)
}
