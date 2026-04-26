package wish

import (
	"context"
	"log"

	errors_entity "github.com/RSODA/wishlist/microservices/wishlist_service/internal/err"
	"github.com/RSODA/wishlist/microservices/wishlist_service/internal/models"
)

func (s *wish) GetWishs(ctx context.Context, req *models.GetWishsRequest) (*models.GetWishsResponse, error) {
	if req == nil || req.TgID <= 0 || req.ToTgID <= 0 {
		return nil, errors_entity.ErrInvalidId
	}

	if req.TgID != req.ToTgID {
		res, err := s.client.GetCheckSub(ctx, req)
		if err != nil {
			log.Println("err check sub: ", err)
			return nil, err
		}

		if !res {
			log.Println("err sub not confirmed: ", res)
			return nil, errors_entity.ErrSubNotConfirmed
		}
	}

	resp, err := s.r.GetWishs(ctx, req)
	if err != nil {
		return nil, err
	}

	for _, v := range resp.Wishs {
		v.Picture = "http://localhost:8081/api/v1/static/" + v.Picture
	}

	return resp, nil
}
