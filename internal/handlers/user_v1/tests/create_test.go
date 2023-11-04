package tests

import (
	"context"
	"errors"
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/gojuno/minimock/v3"
	"github.com/romanfomindev/microservices-auth/internal/convertor"
	handlers "github.com/romanfomindev/microservices-auth/internal/handlers/user_v1"
	"github.com/romanfomindev/microservices-auth/internal/services"
	serviceMock "github.com/romanfomindev/microservices-auth/internal/services/mocks"
	desc "github.com/romanfomindev/microservices-auth/pkg/user_v1"
	"github.com/stretchr/testify/require"
)

func TestCreateHandler(t *testing.T) {
	type userServiceMockFunc func(mc *minimock.Controller) services.UserService

	type args struct {
		ctx     context.Context
		request *desc.CreateRequest
	}

	var (
		ctx = context.Background()
		mc  = minimock.NewController(t)

		id       = gofakeit.Uint64()
		name     = gofakeit.Name()
		email    = gofakeit.Email()
		password = gofakeit.Password(true, false, false, false, false, 32)
		req      = &desc.CreateRequest{
			Info: &desc.UserInfo{
				Name:            name,
				Email:           email,
				Password:        password,
				PasswordConfirm: password,
				Role:            desc.Roles(desc.Roles_value["ADMIN"]),
			},
		}

		res = &desc.CreateResponse{
			Id: id,
		}

		errService = errors.New("service error")
	)

	defer t.Cleanup(mc.Finish)

	tests := []struct {
		name            string
		args            args
		want            *desc.CreateResponse
		err             error
		userServiceMock userServiceMockFunc
	}{
		{
			name: "success case",
			args: args{
				ctx:     ctx,
				request: req,
			},
			want: res,
			err:  nil,
			userServiceMock: func(mc *minimock.Controller) services.UserService {
				mock := serviceMock.NewUserServiceMock(mc)
				mock.CreateMock.Expect(ctx, convertor.ToUserFromDesc(req.GetInfo())).Return(id, nil)

				return mock
			},
		},
		{
			name: "error case",
			args: args{
				ctx:     ctx,
				request: req,
			},
			want: nil,
			err:  errService,
			userServiceMock: func(mc *minimock.Controller) services.UserService {
				mock := serviceMock.NewUserServiceMock(mc)
				mock.CreateMock.Expect(ctx, convertor.ToUserFromDesc(req.GetInfo())).Return(0, errService)

				return mock
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			userServiceMock := tt.userServiceMock(mc)
			handler := handlers.NewUserHandlers(userServiceMock)
			response, err := handler.Create(tt.args.ctx, tt.args.request)
			require.Equal(t, tt.want, response)
			require.Equal(t, tt.err, err)
		})
	}
}
