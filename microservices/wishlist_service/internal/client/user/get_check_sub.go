package user

import (
	"context"
	"log"

	"github.com/RSODA/wishlist/microservices/wishlist_service/internal/models"
	user_v1 "github.com/RSODA/wishlist/pkg/proto/user/v1"
)

func (s *userService) GetCheckSub(ctx context.Context, req *models.GetWishsRequest) (bool, error) {
	res, err := s.UserService.CheckSubIsAccepted(ctx, &user_v1.CheckSubIsAcceptedRequest{
		FromTgId: req.TgID,
		ToTgId:   req.ToTgID,
	})

	log.Println("GetCheckSub res", res)

	if err != nil {
		return false, err
	}

	return res.IsAccepted, nil
}
