package repositories

import (
	"context"

	"github.com/romanfomindev/microservices-auth/internal/models"
	"github.com/romanfomindev/microservices-auth/internal/repositories/user/model"
)

type UserRepository interface {
	Create(ctx context.Context, user model.UserCreate) (uint64, error)
	Update(ctx context.Context, id uint64, user model.UserUpdate) error
	GetById(ctx context.Context, id uint64) (*models.User, error)
	Delete(ctx context.Context, id uint64) error
	GetByEmail(ctx context.Context, email string) (*models.User, error)
}

type UrlsProtectedRepository interface {
	GetByUrl(ctx context.Context, url string) (*models.UrlProtected, error)
}
