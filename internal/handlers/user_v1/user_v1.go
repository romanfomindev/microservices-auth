package handlers

import (
	"context"
	"github.com/brianvoe/gofakeit"
	desc "github.com/romanfomindev/microservices-auth/pkg/user_v1"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"
	"log"
)

type UserV1Service struct {
	desc.UnimplementedUserV1Server
}

func (UserV1Service) Create(ctx context.Context, request *desc.CreateRequest) (*desc.CreateResponse, error) {

	log.Printf("name: %s, email: %s, password: %s, password_confirm: %s, role: %s",
		request.GetInfo().GetName(),
		request.GetInfo().GetEmail(),
		request.GetInfo().GetPassword(),
		request.GetInfo().GetPasswordConfirm(),
		request.GetInfo().GetRole(),
	)

	return &desc.CreateResponse{
		Id: 123,
	}, nil
}

func (UserV1Service) Get(ctx context.Context, request *desc.GetRequest) (*desc.GetResponse, error) {

	log.Printf("ID: %d\n", request.GetId())

	return &desc.GetResponse{
		User: &desc.User{
			Id: request.GetId(),
			Info: &desc.UserInfo{
				Name:            gofakeit.Name(),
				Email:           gofakeit.Email(),
				Password:        "password",
				PasswordConfirm: "password",
				Role:            desc.Roles_USER,
			},
			CreatedAt: timestamppb.New(gofakeit.Date()),
			UpdatedAt: timestamppb.New(gofakeit.Date()),
		},
	}, nil
}

func (UserV1Service) Update(ctx context.Context, request *desc.UpdateRequest) (*emptypb.Empty, error) {
	log.Printf("id: %d, name: %s, email: %s, role: %s",
		request.GetId(),
		request.GetName(),
		request.GetEmail(),
		request.Role,
	)
	return nil, nil
}

func (UserV1Service) Delete(ctx context.Context, request *desc.DeleteRequest) (*emptypb.Empty, error) {

	log.Printf("id: %d", request.GetId())

	return nil, nil
}
