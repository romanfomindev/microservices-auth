package user

import (
	"github.com/romanfomindev/microservices-auth/internal/repositories"
	"github.com/romanfomindev/microservices-auth/internal/services"
)

type service struct {
	repo repositories.UserRepository
}

func NewService(repo repositories.UserRepository) services.UserService {
	return &service{
		repo: repo,
	}
}
