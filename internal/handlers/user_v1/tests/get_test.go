package tests

import (
	"context"
	"errors"
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/gojuno/minimock/v3"
	handlers "github.com/romanfomindev/microservices-auth/internal/handlers/user_v1"
	"github.com/romanfomindev/microservices-auth/internal/logger"
	"github.com/romanfomindev/microservices-auth/internal/models"
	"github.com/romanfomindev/microservices-auth/internal/services"
	serviceMock "github.com/romanfomindev/microservices-auth/internal/services/mocks"
	desc "github.com/romanfomindev/microservices-auth/pkg/user_v1"
	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func setup() {
	logger.Init("info", "testing")
}

func teardown() {

}

func TestGetHandler(t *testing.T) {
	setup()
	defer teardown()

	type userServiceMockFunc func(mc *minimock.Controller) services.UserService

	type args struct {
		ctx     context.Context
		request *desc.GetRequest
	}

	var (
		ctx = context.Background()
		mc  = minimock.NewController(t)

		id        = gofakeit.Uint64()
		name      = gofakeit.Name()
		email     = gofakeit.Email()
		createdAt = gofakeit.Date()
		updatedAt = gofakeit.Date()

		req = &desc.GetRequest{
			Id: id,
		}

		res = &desc.GetResponse{
			User: &desc.User{
				Id: id,
				Info: &desc.UserInfo{
					Name:            name,
					Email:           email,
					Password:        "",
					PasswordConfirm: "",
					Role:            desc.Roles(desc.Roles_value["ADMIN"]),
				},
				CreatedAt: timestamppb.New(createdAt),
				UpdatedAt: timestamppb.New(updatedAt),
			},
		}

		errService = errors.New("service error")

		user = &models.User{
			ID:        int64(id),
			Name:      name,
			Email:     email,
			Role:      models.Role("ADMIN"),
			CreatedAt: createdAt,
			UpdatedAt: &updatedAt,
		}
	)

	defer t.Cleanup(mc.Finish)

	tests := []struct {
		name            string
		args            args
		want            *desc.GetResponse
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
				mock.GetByIdMock.Expect(ctx, id).Return(user, nil)

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
				mock.GetByIdMock.Expect(ctx, id).Return(nil, errService)

				return mock
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			userServiceMock := tt.userServiceMock(mc)
			handler := handlers.NewUserHandlers(userServiceMock)
			response, err := handler.Get(tt.args.ctx, tt.args.request)
			require.Equal(t, tt.want, response)
			require.Equal(t, tt.err, err)
		})
	}
}
