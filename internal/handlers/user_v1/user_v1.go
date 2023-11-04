package handlers

import (
	"context"
	"log"

	"github.com/romanfomindev/microservices-auth/internal/convertor"
	"github.com/romanfomindev/microservices-auth/internal/services"
	desc "github.com/romanfomindev/microservices-auth/pkg/user_v1"
	"google.golang.org/protobuf/types/known/emptypb"
)

type UserV1Handlers struct {
	desc.UnimplementedUserV1Server
	serv services.UserService
}

func NewUserHandlers(service services.UserService) *UserV1Handlers {
	return &UserV1Handlers{
		serv: service,
	}
}

func (h *UserV1Handlers) Create(ctx context.Context, request *desc.CreateRequest) (*desc.CreateResponse, error) {
	log.Printf("name: %s, email: %s, password: %s, password_confirm: %s, role: %s",
		request.GetInfo().GetName(),
		request.GetInfo().GetEmail(),
		request.GetInfo().GetPassword(),
		request.GetInfo().GetPasswordConfirm(),
		request.GetInfo().GetRole(),
	)

	lastInsertId, err := h.serv.Create(
		ctx,
		convertor.ToUserFromDesc(request.GetInfo()),
	)
	if err != nil {
		return nil, err
	}

	return &desc.CreateResponse{
		Id: lastInsertId,
	}, nil
}

func (h *UserV1Handlers) Get(ctx context.Context, request *desc.GetRequest) (*desc.GetResponse, error) {
	log.Printf("ID: %d\n", request.GetId())

	user, err := h.serv.GetById(ctx, request.GetId())
	if err != nil {
		return nil, err
	}

	return convertor.ToUserGetResponseFromUser(user), nil
}

func (h *UserV1Handlers) Update(ctx context.Context, request *desc.UpdateRequest) (*emptypb.Empty, error) {
	log.Printf("id: %d, name: %s, email: %s, role: %s",
		request.GetId(),
		request.GetName(),
		request.GetEmail(),
		request.Role,
	)

	err := h.serv.Update(
		ctx,
		request.GetId(),
		convertor.ToUserFromUpdateRequest(request),
	)
	if err != nil {
		return &emptypb.Empty{}, err
	}

	return &emptypb.Empty{}, nil
}

func (h *UserV1Handlers) Delete(ctx context.Context, request *desc.DeleteRequest) (*emptypb.Empty, error) {
	log.Printf("id: %d", request.GetId())

	err := h.serv.Delete(ctx, request.GetId())

	if err != nil {
		return &emptypb.Empty{}, err
	}

	return &emptypb.Empty{}, nil
}
