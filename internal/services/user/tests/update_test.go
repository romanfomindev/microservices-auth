package tests

import (
	"context"
	"fmt"
	"github.com/brianvoe/gofakeit/v6"
	"github.com/gojuno/minimock/v3"
	"github.com/romanfomindev/microservices-auth/internal/models"
	"github.com/romanfomindev/microservices-auth/internal/repositories"
	repoMock "github.com/romanfomindev/microservices-auth/internal/repositories/mocks"
	"github.com/romanfomindev/microservices-auth/internal/repositories/user/model"
	"github.com/romanfomindev/microservices-auth/internal/services/user"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestUpdate(t *testing.T) {
	type userRepositoryMockFunc func(mc *minimock.Controller) repositories.UserRepository

	type args struct {
		ctx         context.Context
		id          uint64
		userService models.User
	}

	var (
		ctx = context.Background()
		mc  = minimock.NewController(t)

		id      = gofakeit.Uint64()
		name    = gofakeit.Name()
		email   = gofakeit.Email()
		role    = models.Role("ADMIN")
		repoErr = fmt.Errorf("repo error")

		userService = models.User{
			Name:      name,
			Email:     email,
			Password:  "",
			Role:      role,
			CreatedAt: gofakeit.Date(),
		}

		userUpdate = model.UserUpdate{
			Name:  name,
			Email: email,
			Role:  role,
		}

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
		err                error
		userRepositoryMock userRepositoryMockFunc
	}{
		{
			name: "success case",
			args: args{
				ctx:         ctx,
				id:          id,
				userService: userService,
			},
			err: nil,
			userRepositoryMock: func(mc *minimock.Controller) repositories.UserRepository {
				mock := repoMock.NewUserRepositoryMock(mc)
				mock.GetByIdMock.Expect(ctx, id).Return(&res, nil)
				mock.UpdateMock.Expect(ctx, id, userUpdate).Return(nil)
				return mock
			},
		},
		{
			name: "user not found",
			args: args{
				ctx:         ctx,
				id:          id,
				userService: userService,
			},
			err: models.ErrorNoRows,
			userRepositoryMock: func(mc *minimock.Controller) repositories.UserRepository {
				mock := repoMock.NewUserRepositoryMock(mc)
				mock.GetByIdMock.Expect(ctx, id).Return(&res, models.ErrorNoRows)
				return mock
			},
		},
		{
			name: "error database",
			args: args{
				ctx:         ctx,
				id:          id,
				userService: userService,
			},
			err: repoErr,
			userRepositoryMock: func(mc *minimock.Controller) repositories.UserRepository {
				mock := repoMock.NewUserRepositoryMock(mc)
				mock.GetByIdMock.Expect(ctx, id).Return(&res, nil)
				mock.UpdateMock.Expect(ctx, id, userUpdate).Return(repoErr)
				return mock
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			userRepoMock := tt.userRepositoryMock(mc)
			service := user.NewService(userRepoMock)
			err := service.Update(tt.args.ctx, tt.args.id, tt.args.userService)
			require.Equal(t, tt.err, err)
		})
	}
}
