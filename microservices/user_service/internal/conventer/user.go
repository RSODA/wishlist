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

func ToAPIGetSub(items []modelsRepository.Subscribe) []*wishlist.Subscribe {
	var sub []*wishlist.Subscribe

	for _, v := range items {
		sub = append(sub, &wishlist.Subscribe{
			TgId:       v.TgID,
			Username:   v.Username,
			IsAccepted: v.IsAccepted,
		})
	}

	return sub
}

func ToRepoSubscribe(req *modelsService.SubscribeRequest) *modelsRepository.SubscribeRequest {
	return &modelsRepository.SubscribeRequest{
		ToTgID: req.ToTgID,
		TgID:   req.TgID,
	}
}

func ToServiceSubscribe(tgID int64, req *wishlist.SubscribeRequest) *modelsService.SubscribeRequest {
	return &modelsService.SubscribeRequest{
		TgID:       tgID,
		ToUsername: req.ToUsername,
	}
}

func ToRepoAcceptedSub(req *modelsService.AcceptedSubRequest) *modelsRepository.AcceptedSubRequests {
	return &modelsRepository.AcceptedSubRequests{
		TgID:         req.TgID,
		AcceptedTgId: req.AcceptedTgID,
	}
}

func ToServiceAcceptedSub(tgID int64, req *wishlist.AcceptedSubRequest) *modelsService.AcceptedSubRequest {
	return &modelsService.AcceptedSubRequest{
		TgID:         tgID,
		AcceptedTgID: req.AcceptedTgId,
	}
}
