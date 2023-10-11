package main

import (
	"context"
	"github.com/romanfomindev/microservices-auth/internal/config"
	"github.com/romanfomindev/microservices-auth/internal/config/env"
	handlers "github.com/romanfomindev/microservices-auth/internal/handlers/user_v1"
	"github.com/romanfomindev/microservices-auth/internal/managers"
	"github.com/romanfomindev/microservices-auth/internal/repositories/pg"
	desc "github.com/romanfomindev/microservices-auth/pkg/user_v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
)

var configPath string = ".env"

func main() {
	ctx := context.Background()
	// Считываем переменные окружения
	err := config.Load(configPath)
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	grpcConfig, err := env.NewGRPCConfig()
	if err != nil {
		log.Fatalf("failed to get grpc config: %v", err)
	}
	pgConfig, err := env.NewPGConfig()
	if err != nil {
		log.Fatalf("failed to get pg config: %v", err)
	}

	lis, err := net.Listen("tcp", grpcConfig.Address())
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	reflection.Register(s)
	userRepository := pg.NewUserRepository(ctx, pgConfig)
	userManager := managers.NewUserManager(userRepository)
	userService := handlers.NewUserService(userManager)

	desc.RegisterUserV1Server(s, userService)

	if err = s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
