package api

import (
	"context"
	"errors"

	"github.com/RSODA/wishlist/internal/conventer"
	"github.com/RSODA/wishlist/internal/interceptor"
	"github.com/RSODA/wishlist/internal/models"
	wishlist "github.com/RSODA/wishlist/pkg/proto/user/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (i *Implementation) AcceptedSub(ctx context.Context, req *wishlist.AcceptedSubRequest) (*wishlist.AcceptedSubResponse, error) {
	tgID, err := interceptor.TgIDFromContext(ctx)
	if err != nil {
		return nil, err
	}


	err = i.userService.AcceptedSub(ctx, conventer.ToServiceAcceptedSub(tgID, req))
	if err != nil {
		if errors.Is(err, models.ErrUserNotFound) {
			return nil, status.Error(codes.NotFound, "user not found")
		}

		if errors.Is(err, models.ErrInvalidArgument) {
			return nil, status.Error(codes.InvalidArgument, "invalid argument(-s)")
		}

		if errors.Is(err, models.ErrAcceptedSub) {
			return nil, status.Error(codes.Unknown, "err accepted sub")
		}

		return nil, status.Error(codes.Internal, err.Error())
	}

	return nil, nil
}
