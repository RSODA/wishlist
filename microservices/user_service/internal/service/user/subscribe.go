package user

import (
	"context"
	"log"

	"github.com/RSODA/wishlist/internal/conventer"
	"github.com/RSODA/wishlist/internal/models"
)

func (s *userService) Subscribe(ctx context.Context, req *models.SubscribeRequest) error {
	if req.TgID < 0 || req.ToTgID < 0 {
		log.Println("Subscribe: invalid argument", req)
		return models.ErrInvalidArgument
	}

	id, err := s.repo.GetUser(ctx, req.ToUsername)
	if err != nil {
		return err
	}

	req.ToTgID = id

	err = s.repo.Subscribe(ctx, conventer.ToRepoSubscribe(req))
	if err != nil {
		return err
	}

	return nil
}
