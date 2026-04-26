package user

import (
	"context"
	"errors"
	"fmt"

	"github.com/RSODA/wishlist/internal/models"
	"github.com/jackc/pgx/v5"
)

func (s *userService) GetSubIsAccepted(ctx context.Context, req *models.GetSubIsAccepted) (bool, error) {
	if req.ToTgID < 0 || req.FromTgID < 0 {
		return false, fmt.Errorf("invalid param")
	}

	res, err := s.repo.GetSubIsAccepted(ctx, req)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return false, nil
		}

		return false, err
	}

	return res, nil
}
