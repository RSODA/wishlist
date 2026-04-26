package user

import (
	"context"
	"log"

	"github.com/RSODA/wishlist/internal/conventer"
	"github.com/RSODA/wishlist/internal/models"
)

func (s *userService) AcceptedSub(ctx context.Context, req *models.AcceptedSubRequest) error {
	if req.TgID < 0 || req.AcceptedTgID < 0 {
		log.Println("empty params on accepted sub requests")
		return models.ErrInvalidArgument
	}

	err := s.repo.AcceptedSub(ctx, conventer.ToRepoAcceptedSub(req))
	if err != nil {
		return err
	}

	return nil
}
