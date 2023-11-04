package tests

import (
	"context"
	"fmt"
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/gojuno/minimock/v3"
	"github.com/romanfomindev/microservices-auth/internal/repositories"
	repoMock "github.com/romanfomindev/microservices-auth/internal/repositories/mocks"
	"github.com/romanfomindev/microservices-auth/internal/services/user"
	"github.com/stretchr/testify/require"
)

func TestDelete(t *testing.T) {
	type userRepositoryMockFunc func(mc *minimock.Controller) repositories.UserRepository

	type args struct {
		ctx context.Context
		id  uint64
	}

	var (
		ctx = context.Background()
		mc  = minimock.NewController(t)

		id      = gofakeit.Uint64()
		repoErr = fmt.Errorf("repo error")
	)

	defer t.Cleanup(mc.Finish)

	tests := []struct {
		name               string
		args               args
		err                error
		userRepositoryMock userRepositoryMockFunc
	}{
		{
			name: "case success",
			args: args{
				ctx: ctx,
				id:  id,
			},
			err: nil,
			userRepositoryMock: func(mc *minimock.Controller) repositories.UserRepository {
				mock := repoMock.NewUserRepositoryMock(mc)
				mock.DeleteMock.Expect(ctx, id).Return(nil)

				return mock
			},
		},
		{
			name: "case error",
			args: args{
				ctx: ctx,
				id:  id,
			},
			err: repoErr,
			userRepositoryMock: func(mc *minimock.Controller) repositories.UserRepository {
				mock := repoMock.NewUserRepositoryMock(mc)
				mock.DeleteMock.Expect(ctx, id).Return(repoErr)

				return mock
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			userRepoMock := tt.userRepositoryMock(mc)
			service := user.NewService(userRepoMock)
			err := service.Delete(tt.args.ctx, tt.args.id)
			require.Equal(t, tt.err, err)
		})
	}
}
