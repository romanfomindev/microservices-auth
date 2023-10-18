package convertor

import (
	"github.com/romanfomindev/microservices-auth/internal/models"
	"github.com/romanfomindev/microservices-auth/internal/repositories/user/model"
	desc "github.com/romanfomindev/microservices-auth/pkg/user_v1"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func ToUserCreateFromUserService(userService models.User) model.UserCreate {
	return model.UserCreate{
		Name:     userService.Name,
		Email:    userService.Email,
		Password: userService.Password,
		Role:     userService.Role,
	}
}

func ToUserUpdateFromUserService(userService models.User) model.UserUpdate {
	return model.UserUpdate{
		Name:  userService.Name,
		Email: userService.Email,
		Role:  userService.Role,
	}
}

func ToUserFromDesc(userInfo *desc.UserInfo) models.User {
	return models.User{
		Name:     userInfo.GetName(),
		Email:    userInfo.GetEmail(),
		Password: userInfo.GetPassword(),
		Role:     models.Role(userInfo.GetRole().String()),
	}
}

func ToUserFromUpdateRequest(userInfo *desc.UpdateRequest) models.User {
	return models.User{
		Name:  userInfo.GetName().GetValue(),
		Email: userInfo.GetEmail().GetValue(),
		Role:  models.Role(userInfo.GetRole().String()),
	}
}

func ToUserGetResponseFromUser(userService *models.User) *desc.GetResponse {
	var updatedAt *timestamppb.Timestamp
	if userService.UpdatedAt != nil {
		updatedAt = timestamppb.New(*userService.UpdatedAt)
	}
	return &desc.GetResponse{
		User: &desc.User{
			Id: uint64(userService.ID),
			Info: &desc.UserInfo{
				Name:  userService.Name,
				Email: userService.Email,
				Role:  desc.Roles(desc.Roles_value[string(userService.Role)]),
			},
			CreatedAt: timestamppb.New(userService.CreatedAt),
			UpdatedAt: updatedAt,
		},
	}
}
