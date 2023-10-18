package user

import (
	"context"
	"time"

	"github.com/romanfomindev/microservices-auth/internal/convertor"
	"github.com/romanfomindev/microservices-auth/internal/models"
)

func (m *service) Update(ctx context.Context, id uint64, userService models.User) error {
	ctx, cancel := context.WithTimeout(ctx, time.Second*10)
	defer func() {
		cancel()
	}()
	userRepo := convertor.ToUserUpdateFromUserService(userService)
	return m.repo.Update(ctx, id, userRepo)
}
