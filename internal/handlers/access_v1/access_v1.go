package access_v1

import (
	"context"
	"errors"
	"strings"

	"github.com/romanfomindev/microservices-auth/internal/services"
	desc "github.com/romanfomindev/microservices-auth/pkg/access_v1"
	"google.golang.org/grpc/metadata"
	"google.golang.org/protobuf/types/known/emptypb"
)

const (
	authPrefix = "Bearer "
)

type AccessV1Handlers struct {
	desc.UnimplementedAccessServiceServer
	service services.AccessService
}

func NewAccessHandlers(service services.AccessService) *AccessV1Handlers {
	return &AccessV1Handlers{
		service: service,
	}
}

func (h *AccessV1Handlers) Check(ctx context.Context, request *desc.CheckRequest) (*emptypb.Empty, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, errors.New("metadata is not provided")
	}

	authHeader, ok := md["authorization"]
	if !ok || len(authHeader) == 0 {
		return nil, errors.New("authorization header is not provided")
	}

	if !strings.HasPrefix(authHeader[0], authPrefix) {
		return nil, errors.New("invalid authorization header format")
	}

	accessToken := strings.TrimPrefix(authHeader[0], authPrefix)

	err := h.service.Check(ctx, accessToken, request.EndpointAddress)
	if err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}
