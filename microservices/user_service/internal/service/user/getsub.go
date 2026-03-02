package user

import (
	"context"
	"log"

	"github.com/RSODA/wishlist/internal/models"
)

func (s *userService) GetSub(ctx context.Context, tgId int64) (*[]models.Subscribe, error) {
	if tgId < 0 {
		log.Println("invalid tgId", tgId)
		return nil, models.ErrInvalidId
	}

	res, err := s.repo.GetSub(ctx, tgId)
	if err != nil {
		return nil, err
	}

	return res, nil
}
