package user

import (
	"context"
	"time"

	"github.com/romanfomindev/microservices-auth/internal/models"
)

func (m *service) GetById(ctx context.Context, id uint64) (*models.User, error) {
	ctx, cancel := context.WithTimeout(ctx, time.Second*10)
	defer func() {
		cancel()
	}()

	return m.repo.GetById(ctx, id)
}
