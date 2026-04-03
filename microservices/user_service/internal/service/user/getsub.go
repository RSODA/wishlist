package user

import (
	"context"
	"fmt"
	"log"

	customErr "github.com/RSODA/wishlist/internal/models"
	"github.com/RSODA/wishlist/internal/repository/models"
)

func (s *userService) GetSub(ctx context.Context, tgId int64, isAccept bool) (*models.GetSubResponse, error) {

	if tgId < 0 {
		log.Println("invalid tgId", tgId)
		return nil, customErr.ErrInvalidId
	}

	res, err := s.repo.GetSub(ctx, tgId, isAccept)
	fmt.Println(res)
	if err != nil {
		return nil, err
	}
	return res, nil
}
