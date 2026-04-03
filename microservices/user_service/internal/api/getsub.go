package api

import (
	"context"
	"errors"
	"fmt"

	"github.com/RSODA/wishlist/internal/conventer"
	"github.com/RSODA/wishlist/internal/models"
	wishlist "github.com/RSODA/wishlist/pkg/proto/user/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (i *Implementation) GetSub(ctx context.Context, req *wishlist.GetSubRequest) (*wishlist.GetSubResponse, error) {
	res, err := i.userService.GetSub(ctx, req.TgId, req.IsAccepted)

	fmt.Println(res)
	if err != nil {
		if errors.Is(err, models.ErrInvalidId) {
			return nil, status.Error(codes.InvalidArgument, "invalid user id")
		}

		if errors.Is(err, models.ErrUserNotFound) {
			return nil, status.Error(codes.NotFound, "user not found")
		}

		if errors.Is(err, models.ErrGetSub) {
			return nil, status.Error(codes.Unknown, "get sub err")
		}

		return nil, status.Error(codes.Internal, "db error")
	}

	fmt.Println(&wishlist.GetSubResponse{Username: res.Username, Follower: conventer.ToAPIGetSub(res.Sub)})

	return &wishlist.GetSubResponse{Username: res.Username, Follower: conventer.ToAPIGetSub(res.Sub)}, nil
}
