package user

import (
	"context"
)

func (s *service) Delete(ctx context.Context, id uint64) error {
	return s.repo.Delete(ctx, id)
}
