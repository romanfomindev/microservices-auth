package convertor

import (
	"time"

	"github.com/romanfomindev/microservices-auth/internal/models"
	"github.com/romanfomindev/microservices-auth/internal/repositories/user/model"
)

func ToUserFromUserRepo(userRepo model.User) *models.User {
	var updatedAt *time.Time
	if userRepo.UpdatedAt.Valid {
		updatedAt = &userRepo.UpdatedAt.Time
	}

	return &models.User{
		ID:        userRepo.ID,
		Name:      userRepo.Name,
		Email:     userRepo.Email,
		Role:      userRepo.Role,
		CreatedAt: userRepo.CreatedAt,
		UpdatedAt: updatedAt,
	}
}
