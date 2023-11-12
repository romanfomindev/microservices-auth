package access

import (
	"github.com/romanfomindev/microservices-auth/internal/repositories"
	"github.com/romanfomindev/microservices-auth/internal/services"
)

type accessService struct {
	repo                 repositories.UrlsProtectedRepository
	accessTokenSecretKey string
}

func NewAccessService(repo repositories.UrlsProtectedRepository, accessTokenSecretKey string) services.AccessService {
	return &accessService{
		repo:                 repo,
		accessTokenSecretKey: accessTokenSecretKey,
	}
}
