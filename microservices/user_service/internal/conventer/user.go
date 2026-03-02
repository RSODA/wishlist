package conventer

import (
	modelsService "github.com/RSODA/wishlist/internal/models"
	modelsRepository "github.com/RSODA/wishlist/internal/repository/models"
	wishlist "github.com/RSODA/wishlist/pkg/proto/user/v1"
)

func ToRepoUserCreate(userService *modelsService.CreateUserRequest) *modelsRepository.CreateUserRequest {
	userCreateRepo := &modelsRepository.CreateUserRequest{
		Username: userService.Username,
		TgID:     userService.TgID,
	}

	return userCreateRepo
}

func ToServiceUserCreate(userCreateAPI *wishlist.CreateUserRequest) *modelsService.CreateUserRequest {
	return &modelsService.CreateUserRequest{
		Username: userCreateAPI.Username,
		TgID:     userCreateAPI.TgId,
	}
}
