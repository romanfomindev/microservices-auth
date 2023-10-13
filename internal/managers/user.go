package managers

import (
	"context"
	"time"

	"github.com/romanfomindev/microservices-auth/internal/models"
	"github.com/romanfomindev/microservices-auth/internal/repositories"
)

type UserManager struct {
	repo repositories.UserRepository
}

func NewUserManager(repo repositories.UserRepository) *UserManager {
	return &UserManager{
		repo: repo,
	}
}

func (m *UserManager) Create(ctx context.Context, name, email, password, role string) (uint64, error) {
	ctx, cancel := context.WithTimeout(ctx, time.Second*10)
	defer func() {
		cancel()
	}()

	return m.repo.Create(ctx, name, email, password, role)
}

func (m *UserManager) Update(ctx context.Context, id uint64, name, email, role string) error {
	ctx, cancel := context.WithTimeout(ctx, time.Second*10)
	defer func() {
		cancel()
	}()

	return m.repo.Update(ctx, id, name, email, role)
}

func (m *UserManager) GetById(ctx context.Context, id uint64) (models.User, error) {
	ctx, cancel := context.WithTimeout(ctx, time.Second*10)
	defer func() {
		cancel()
	}()

	return m.repo.GetById(ctx, id)
}

func (m *UserManager) Delete(ctx context.Context, id uint64) error {
	ctx, cancel := context.WithTimeout(ctx, time.Second*10)
	defer func() {
		cancel()
	}()

	return m.repo.Delete(ctx, id)
}
