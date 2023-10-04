package main

import (
	"fmt"
	handlers "github.com/romanfomindev/microservices-auth/internal/handlers/user_v1"
	desc "github.com/romanfomindev/microservices-auth/pkg/user_v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
)

const grpcPort = 50051

func main() {

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", grpcPort))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	reflection.Register(s)
	userService := handlers.UserV1Service{}
	desc.RegisterUserV1Server(s, userService)

	if err = s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
