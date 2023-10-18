package services

import (
	"context"

	"github.com/romanfomindev/microservices-auth/internal/models"
)

type UserService interface {
	Create(ctx context.Context, userService models.User) (uint64, error)
	Update(ctx context.Context, id uint64, userService models.User) error
	Delete(ctx context.Context, id uint64) error
	GetById(ctx context.Context, id uint64) (*models.User, error)
}
