package user

import (
	"context"
	"time"

	"github.com/romanfomindev/microservices-auth/internal/convertor"
	"github.com/romanfomindev/microservices-auth/internal/models"
)

func (m *service) Create(ctx context.Context, userService models.User) (uint64, error) {
	ctx, cancel := context.WithTimeout(ctx, time.Second*10)
	defer func() {
		cancel()
	}()
	userRepo := convertor.ToUserCreateFromUserService(userService)
	return m.repo.Create(ctx, userRepo)
}
