package repositories

import (
	"context"

	"github.com/romanfomindev/microservices-auth/internal/models"
)

type UserRepository interface {
	Create(ctx context.Context, name, email, password, role string) (uint64, error)
	Update(ctx context.Context, id uint64, name, email, role string) error
	GetById(ctx context.Context, id uint64) (*models.User, error)
	Delete(ctx context.Context, id uint64) error
}
