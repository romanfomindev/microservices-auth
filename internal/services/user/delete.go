package user

import (
	"context"
	"time"
)

func (m *service) Delete(ctx context.Context, id uint64) error {
	ctx, cancel := context.WithTimeout(ctx, time.Second*10)
	defer func() {
		cancel()
	}()

	return m.repo.Delete(ctx, id)
}
