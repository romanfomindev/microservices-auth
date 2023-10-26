package tests

import (
	"context"
	"fmt"
	"github.com/brianvoe/gofakeit/v6"
	"github.com/gojuno/minimock/v3"
	"github.com/romanfomindev/microservices-auth/internal/models"
	"github.com/romanfomindev/microservices-auth/internal/repositories"
	repoMock "github.com/romanfomindev/microservices-auth/internal/repositories/mocks"
	"github.com/romanfomindev/microservices-auth/internal/services/user"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestGet(t *testing.T) {
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

		res = models.User{
			ID:        gofakeit.Int64(),
			Name:      gofakeit.Name(),
			Email:     gofakeit.Email(),
			Password:  "",
			Role:      models.Role("ADMIN"),
			CreatedAt: gofakeit.Date(),
		}
	)

	defer t.Cleanup(mc.Finish)

	tests := []struct {
		name               string
		args               args
		want               *models.User
		err                error
		userRepositoryMock userRepositoryMockFunc
	}{
		{
			name: "success case",
			args: args{
				ctx: ctx,
				id:  id,
			},
			want: &res,
			err:  nil,
			userRepositoryMock: func(mc *minimock.Controller) repositories.UserRepository {
				mock := repoMock.NewUserRepositoryMock(mc)
				mock.GetByIdMock.Expect(ctx, id).Return(&res, nil)

				return mock
			},
		},
		{
			name: "error repo case",
			args: args{
				ctx: ctx,
				id:  id,
			},
			want: nil,
			err:  repoErr,
			userRepositoryMock: func(mc *minimock.Controller) repositories.UserRepository {
				mock := repoMock.NewUserRepositoryMock(mc)
				mock.GetByIdMock.Expect(ctx, id).Return(nil, repoErr)

				return mock
			},
		},
		{
			name: "error now row",
			args: args{
				ctx: ctx,
				id:  id,
			},
			want: nil,
			err:  models.ErrorNoRows,
			userRepositoryMock: func(mc *minimock.Controller) repositories.UserRepository {
				mock := repoMock.NewUserRepositoryMock(mc)
				mock.GetByIdMock.Expect(ctx, id).Return(nil, models.ErrorNoRows)

				return mock
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			userRepoMock := tt.userRepositoryMock(mc)
			service := user.NewService(userRepoMock)
			newId, err := service.GetById(tt.args.ctx, tt.args.id)
			require.Equal(t, tt.want, newId)
			require.Equal(t, tt.err, err)
		})
	}
}
