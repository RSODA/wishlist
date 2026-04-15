package wish

import (
	"context"
	"log"

	"github.com/RSODA/wishlist/microservices/wishlist_service/internal/err"
	"github.com/RSODA/wishlist/microservices/wishlist_service/internal/media"
	"github.com/RSODA/wishlist/microservices/wishlist_service/internal/models"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *wish) Create(ctx context.Context, req *models.CreateRequest) (*int64, error) {
	if err := s.validate.Struct(req); err != nil {
		log.Println("err validate struct CreateRequest: ", err)
		return nil, errors_entity.ErrValidateWish
	}

	r, err := s.client.GetUserSub(ctx, req.TgID)
	if err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, errors_entity.ErrUserNotFound
		}
		return nil, err
	}

	if r == nil || r.Username == "" {
		return nil, errors_entity.ErrUserNotFound
	}

	picture, err := media.Upload(req.Picture)
	if err != nil {
		return nil, err
	}

	req.Picture = picture

	resp, err := s.r.Create(ctx, req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
