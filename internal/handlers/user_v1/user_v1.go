package handlers

import (
	"context"
	"github.com/brianvoe/gofakeit"
	"github.com/romanfomindev/microservices-auth/internal/managers"
	"github.com/romanfomindev/microservices-auth/internal/models"
	desc "github.com/romanfomindev/microservices-auth/pkg/user_v1"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"
	"log"
)

type UserV1Service struct {
	desc.UnimplementedUserV1Server
	manager *managers.UserManager
}

func NewUserService(manager *managers.UserManager) desc.UserV1Server {
	return &UserV1Service{
		manager: manager,
	}
}

func (s *UserV1Service) Create(ctx context.Context, request *desc.CreateRequest) (*desc.CreateResponse, error) {
	log.Printf("name: %s, email: %s, password: %s, password_confirm: %s, role: %s",
		request.GetInfo().GetName(),
		request.GetInfo().GetEmail(),
		request.GetInfo().GetPassword(),
		request.GetInfo().GetPasswordConfirm(),
		request.GetInfo().GetRole(),
	)

	lastInsertId, err := s.manager.Create(
		ctx,
		request.GetInfo().GetName(),
		request.GetInfo().GetEmail(),
		request.GetInfo().GetPassword(),
		models.ROLE_USER,
	)

	return &desc.CreateResponse{
		Id: lastInsertId,
	}, err
}

func (s *UserV1Service) Get(ctx context.Context, request *desc.GetRequest) (*desc.GetResponse, error) {
	log.Printf("ID: %d\n", request.GetId())

	user := s.manager.GetById(ctx, request.GetId())

	/** TODO трансформатор с grpc to json */
	return &desc.GetResponse{
		User: &desc.User{
			Id: request.GetId(),
			Info: &desc.UserInfo{
				Name:  user.Name,
				Email: user.Email,
				Role:  desc.Roles_USER,
			},
			CreatedAt: timestamppb.New(gofakeit.Date()),
			UpdatedAt: timestamppb.New(gofakeit.Date()),
		},
	}, nil
}

func (s *UserV1Service) Update(ctx context.Context, request *desc.UpdateRequest) (*emptypb.Empty, error) {
	log.Printf("id: %d, name: %s, email: %s, role: %s",
		request.GetId(),
		request.GetName(),
		request.GetEmail(),
		request.Role,
	)

	err := s.manager.Update(ctx, request.GetId(), request.GetName().GetValue(), request.GetEmail().GetValue(), models.ROLE_USER)

	if err != nil {
		return &emptypb.Empty{}, err
	}

	return &emptypb.Empty{}, nil
}

func (s *UserV1Service) Delete(ctx context.Context, request *desc.DeleteRequest) (*emptypb.Empty, error) {
	log.Printf("id: %d", request.GetId())

	err := s.manager.Delete(ctx, request.GetId())

	if err != nil {
		return &emptypb.Empty{}, err
	}

	return &emptypb.Empty{}, nil
}
