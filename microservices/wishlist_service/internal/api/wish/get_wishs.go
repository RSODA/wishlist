package wish

import (
	"context"
	"errors"

	"github.com/RSODA/wishlist/microservices/wishlist_service/internal/conventer"
	errors_entity "github.com/RSODA/wishlist/microservices/wishlist_service/internal/err"
	"github.com/RSODA/wishlist/microservices/wishlist_service/internal/models"
	wish_v1 "github.com/RSODA/wishlist/microservices/wishlist_service/pkg/proto/wish/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (i *Implementation) GetWishs(ctx context.Context, req *wish_v1.GetWishsRequest) (*wish_v1.GetWishsResponse, error) {
	resp, err := i.wishService.GetWishs(ctx, &models.GetWishsRequest{
		TgID:   req.TgId,
		Offset: req.Offset,
	})

	if err != nil {
		if errors.Is(err, errors_entity.ErrInvalidId) {
			return nil, status.Error(codes.InvalidArgument, err.Error())
		}

		return nil, status.Error(codes.Internal, err.Error())
	}

	return conventer.ToAPIFromService(resp), nil
}
