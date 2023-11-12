package access

import (
	"context"
	"errors"

	"github.com/romanfomindev/microservices-auth/internal/models"
	"github.com/romanfomindev/microservices-auth/internal/utils"
)

func (s accessService) Check(ctx context.Context, accessToken, endpointAddress string) error {
	claims, err := utils.VerifyToken(accessToken, []byte(s.accessTokenSecretKey))
	if err != nil {
		return errors.New("access token is invalid: " + err.Error())
	}

	urlProtected, err := s.repo.GetByUrl(ctx, endpointAddress)
	if err != nil {
		if errors.Is(err, models.ErrorNoRows) {
			return nil
		}
		return err
	}

	for _, role := range urlProtected.Roles {
		if role == claims.Role {
			return nil
		}
	}

	return models.ErrorAccessDenied
}
