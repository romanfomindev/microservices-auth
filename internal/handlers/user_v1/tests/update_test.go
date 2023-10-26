package tests

import (
	"context"
	"errors"
	handlers "github.com/romanfomindev/microservices-auth/internal/handlers/user_v1"
	"github.com/romanfomindev/microservices-auth/internal/models"
	serviceMock "github.com/romanfomindev/microservices-auth/internal/services/mocks"
	"github.com/stretchr/testify/require"
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/gojuno/minimock/v3"
	"github.com/romanfomindev/microservices-auth/internal/services"
	desc "github.com/romanfomindev/microservices-auth/pkg/user_v1"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

func TestUpdateHandler(t *testing.T) {
	type userServiceMockFunc func(mc *minimock.Controller) services.UserService

	type args struct {
		ctx     context.Context
		request *desc.UpdateRequest
	}

	var (
		ctx = context.Background()
		mc  = minimock.NewController(t)

		id    = gofakeit.Uint64()
		name  = gofakeit.Name()
		email = gofakeit.Email()
		role  = desc.Roles(desc.Roles_value["ADMIN"])

		req = &desc.UpdateRequest{
			Id:    id,
			Name:  wrapperspb.String(name),
			Email: wrapperspb.String(email),
			Role:  role,
		}
		res = &emptypb.Empty{}
		errService = errors.New("service error")

		user = &models.User{
			Name:      name,
			Email:     email,
			Role:      models.Role("ADMIN"),
		}
	)
	defer t.Cleanup(mc.Finish)

	tests := []struct {
		name            string
		args            args
		want            *emptypb.Empty
		err             error
		userServiceMock userServiceMockFunc
	}{
		{
			name: "success case",
			args: args{
				ctx: ctx,
				request: req,
			},
			want: res,
			err: nil,
			userServiceMock: func(mc *minimock.Controller) services.UserService {
				mock := serviceMock.NewUserServiceMock(mc)
				mock.UpdateMock.Expect(ctx, id, *user).Return(nil)

				return mock
			},
		},
		{
			name: "error case user not found",
			args: args{
				ctx: ctx,
				request: req,
			},
			want: res,
			err: models.ErrorNoRows,
			userServiceMock: func(mc *minimock.Controller) services.UserService {
				mock := serviceMock.NewUserServiceMock(mc)
				mock.UpdateMock.Expect(ctx, id, *user).Return(models.ErrorNoRows)

				return mock
			},
		},
		{
			name: "error case",
			args: args{
				ctx: ctx,
				request: req,
			},
			want: res,
			err: errService,
			userServiceMock: func(mc *minimock.Controller) services.UserService {
				mock := serviceMock.NewUserServiceMock(mc)
				mock.UpdateMock.Expect(ctx, id, *user).Return(errService)

				return mock
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			userServiceMock := tt.userServiceMock(mc)
			handler := handlers.NewUserHandlers(userServiceMock)
			response, err := handler.Update(tt.args.ctx, tt.args.request)
			require.Equal(t, tt.want, response)
			require.Equal(t, tt.err, err)
		})
	}
}
