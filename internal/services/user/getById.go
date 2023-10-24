package user

import (
	"context"
	"github.com/romanfomindev/microservices-auth/internal/models"
)

func (s *service) GetById(ctx context.Context, id uint64) (*models.User, error) {
	return s.repo.GetById(ctx, id)
}
