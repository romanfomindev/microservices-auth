package tests

import (
	"context"
	"errors"
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/gojuno/minimock/v3"
	handlers "github.com/romanfomindev/microservices-auth/internal/handlers/user_v1"
	"github.com/romanfomindev/microservices-auth/internal/services"
	serviceMock "github.com/romanfomindev/microservices-auth/internal/services/mocks"
	desc "github.com/romanfomindev/microservices-auth/pkg/user_v1"
	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/types/known/emptypb"
)

func TestDeleteHandler(t *testing.T) {
	type userServiceMockFunc func(mc *minimock.Controller) services.UserService

	type args struct {
		ctx     context.Context
		request *desc.DeleteRequest
	}

	var (
		ctx = context.Background()
		mc  = minimock.NewController(t)

		id = gofakeit.Uint64()

		req = &desc.DeleteRequest{
			Id: id,
		}

		res = &emptypb.Empty{}

		errService = errors.New("service error")
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
				ctx:     ctx,
				request: req,
			},
			want: res,
			err:  nil,
			userServiceMock: func(mc *minimock.Controller) services.UserService {
				mock := serviceMock.NewUserServiceMock(mc)
				mock.DeleteMock.Expect(ctx, id).Return(nil)

				return mock
			},
		},
		{
			name: "error case",
			args: args{
				ctx:     ctx,
				request: req,
			},
			want: res,
			err:  errService,
			userServiceMock: func(mc *minimock.Controller) services.UserService {
				mock := serviceMock.NewUserServiceMock(mc)
				mock.DeleteMock.Expect(ctx, id).Return(errService)

				return mock
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			userServiceMock := tt.userServiceMock(mc)
			handler := handlers.NewUserHandlers(userServiceMock)
			response, err := handler.Delete(tt.args.ctx, tt.args.request)
			require.Equal(t, tt.want, response)
			require.Equal(t, tt.err, err)
		})
	}
}
