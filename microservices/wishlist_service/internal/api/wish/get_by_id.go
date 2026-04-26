package wish

import (
	"context"
	"errors"

	errors_entity "github.com/RSODA/wishlist/microservices/wishlist_service/internal/err"
	wish_v1 "github.com/RSODA/wishlist/microservices/wishlist_service/pkg/proto/wish/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (i *Implementation) GetWishById(ctx context.Context, request *wish_v1.GetWishByIdRequest) (*wish_v1.GetWishByIdResponse, error) {

	res, err := i.wishService.GetById(ctx, request.Id)
	if err != nil {
		if errors.Is(err, errors_entity.ErrInvalidId) {
			return nil, status.Error(codes.InvalidArgument, err.Error())
		}

		if errors.Is(err, errors_entity.ErrGetWishById) {
			return nil, status.Error(codes.InvalidArgument, err.Error())
		}

		return nil, status.Error(codes.NotFound, err.Error())
	}

	return &wish_v1.GetWishByIdResponse{
		Wish: &wish_v1.Wish{
			Id:      res.ID,
			TgId:    res.TgID,
			Title:   res.Title,
			Price:   res.Price,
			Url:     res.URL,
			Picture: res.Picture,
		},
	}, nil
}
