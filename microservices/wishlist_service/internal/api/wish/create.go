package wish

import (
	"context"
	"errors"
	"log"

	"github.com/RSODA/wishlist/microservices/wishlist_service/internal/err"
	"github.com/RSODA/wishlist/microservices/wishlist_service/internal/models"
	wish_v1 "github.com/RSODA/wishlist/microservices/wishlist_service/pkg/proto/wish/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (i *Implementation) CreateWish(ctx context.Context, request *wish_v1.CreateWishRequest) (*wish_v1.CreateWishResponse, error) {
	var resp wish_v1.CreateWishResponse

	id, err := i.wishService.Create(ctx, &models.CreateRequest{
		TgID:    request.TgId,
		Title:   request.Title,
		Price:   request.Price,
		URL:     request.Url,
		Picture: request.Picture,
	})

	if err != nil {
		if errors.Is(err, errors_entity.ErrUserNotFound) {
			return nil, status.Error(codes.NotFound, "user not found")
		}

		if errors.Is(err, errors_entity.ErrValidateWish) {
			return nil, status.Error(codes.InvalidArgument, "err invalid argument")
		}

		if errors.Is(err, errors_entity.ErrCreateWish) {
			return nil, status.Error(codes.Unknown, "err service")
		}

		return nil, status.Error(codes.Unknown, "err service")
	}

	resp.Id = *id

	log.Println("successful created wishlist")
	return &resp, nil
}
