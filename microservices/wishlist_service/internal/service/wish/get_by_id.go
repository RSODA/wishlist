package wish

import (
	"context"

	errors_entity "github.com/RSODA/wishlist/microservices/wishlist_service/internal/err"
	"github.com/RSODA/wishlist/microservices/wishlist_service/internal/media"
	"github.com/RSODA/wishlist/microservices/wishlist_service/internal/models"
)

func (s *wish) GetById(ctx context.Context, id int64) (*models.Wish, error) {
	if id < 0 {
		return nil, errors_entity.ErrInvalidId
	}

	res, err := s.r.GetById(ctx, id)
	if err != nil {
		return nil, err
	}

	res.Picture = media.StaticUrl(s.host, res.Picture)

	return res, nil
}
