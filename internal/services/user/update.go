package user

import (
	"context"
	"github.com/romanfomindev/microservices-auth/internal/convertor"
	"github.com/romanfomindev/microservices-auth/internal/models"
)

func (s *service) Update(ctx context.Context, id uint64, userService models.User) error {
	_, err := s.repo.GetById(ctx, id)
	if err != nil {
		return err
	}

	userRepo := convertor.ToUserUpdateFromUserService(userService)
	return s.repo.Update(ctx, id, userRepo)
}
