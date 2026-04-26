package user

import (
	"context"
	"log"

	customErr "github.com/RSODA/wishlist/internal/models"
	"github.com/RSODA/wishlist/internal/repository/models"
)

func (s *userService) GetSub(ctx context.Context, tgId int64) (*models.GetSubResponse, error) {

	if tgId < 0 {
		log.Println("invalid tgId", tgId)
		return nil, customErr.ErrInvalidId
	}

	res, err := s.repo.GetSubFrom(ctx, tgId)
	if err != nil {
		return nil, err
	}

	st, err := s.repo.GetSubTo(ctx, tgId)
	if err != nil {
		return nil, err
	}

	res.SubTo = st

	return res, nil
}
