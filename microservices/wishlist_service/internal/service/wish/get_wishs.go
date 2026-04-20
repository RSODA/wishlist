package wish

import (
	"context"

	errors_entity "github.com/RSODA/wishlist/microservices/wishlist_service/internal/err"
	"github.com/RSODA/wishlist/microservices/wishlist_service/internal/models"
)

func (s *wish) GetWishs(ctx context.Context, req *models.GetWishsRequest) (*models.GetWishsResponse, error) {
	if req.TgID < 0 {
		return nil, errors_entity.ErrInvalidId
	}

	resp, err := s.r.GetWishs(ctx, req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
