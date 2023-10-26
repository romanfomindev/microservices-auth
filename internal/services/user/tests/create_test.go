package tests

import (
	"context"
	"fmt"
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/gojuno/minimock/v3"
	"github.com/romanfomindev/microservices-auth/internal/models"
	"github.com/romanfomindev/microservices-auth/internal/repositories"
	repoMock "github.com/romanfomindev/microservices-auth/internal/repositories/mocks"
	"github.com/romanfomindev/microservices-auth/internal/repositories/user/model"
	"github.com/romanfomindev/microservices-auth/internal/services/user"
	"github.com/stretchr/testify/require"
)

func TestCreate(t *testing.T) {
	type userRepositoryMockFunc func(mc *minimock.Controller) repositories.UserRepository

	type args struct {
		ctx context.Context
		req models.User
	}

	var (
		ctx = context.Background()
		mc  = minimock.NewController(t)

		id       = gofakeit.Uint64()
		name     = gofakeit.Name()
		email    = gofakeit.Email()
		password = gofakeit.Password(true, false, false, false, false, 32)
		role     = models.Role("ADMIN")

		repoErr = fmt.Errorf("repo error")

		req = models.User{
			Name:     name,
			Email:    email,
			Password: password,
			Role:     role,
		}
		reqRepo = model.UserCreate{
			Name:     name,
			Email:    email,
			Password: password,
			Role:     role,
		}
	)

	defer t.Cleanup(mc.Finish)

	tests := []struct {
		name               string
		args               args
		want               uint64
		err                error
		userRepositoryMock userRepositoryMockFunc
	}{
		{
			name: "success case",
			args: args{
				ctx: ctx,
				req: req,
			},
			want: id,
			err:  nil,
			userRepositoryMock: func(mc *minimock.Controller) repositories.UserRepository {
				mock := repoMock.NewUserRepositoryMock(mc)
				mock.CreateMock.Expect(ctx, reqRepo).Return(id, nil)

				return mock
			},
		},
		{
			name: "error case",
			args: args{
				ctx: ctx,
				req: req,
			},
			want: 0,
			err:  repoErr,
			userRepositoryMock: func(mc *minimock.Controller) repositories.UserRepository {
				mock := repoMock.NewUserRepositoryMock(mc)
				mock.CreateMock.Expect(ctx, reqRepo).Return(0, repoErr)

				return mock
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			userRepoMock := tt.userRepositoryMock(mc)
			service := user.NewService(userRepoMock)
			newId, err := service.Create(tt.args.ctx, tt.args.req)
			require.Equal(t, tt.want, newId)
			require.Equal(t, tt.err, err)
		})
	}
}
