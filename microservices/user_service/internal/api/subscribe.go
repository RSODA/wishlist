package api

import (
	"context"
	"errors"
	"strconv"
	"strings"

	"github.com/RSODA/wishlist/internal/conventer"
	"github.com/RSODA/wishlist/internal/models"
	wishlist "github.com/RSODA/wishlist/pkg/proto/user/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

func (i *Implementation) Subscribe(ctx context.Context, req *wishlist.SubscribeRequest) (*wishlist.SubscribeResponse, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, status.Error(codes.InvalidArgument, "missing metadata")
	}

	mdAuth := md.Get("authorization")
	if len(mdAuth) == 0 {
		return nil, status.Error(codes.InvalidArgument, "missing token")
	}

	const prefix = "Bearer "
	if !strings.HasPrefix(mdAuth[0], prefix) {
		return nil, status.Error(codes.Unauthenticated, "invalid token format")
	}

	tokenStr := strings.TrimPrefix(mdAuth[0], prefix)

	tgId, err := strconv.ParseInt(tokenStr, 10, 64)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "invalid parse token")
	}

	err = i.userService.Subscribe(ctx, conventer.ToServiceSubscribe(tgId, req))
	if err != nil {
		if errors.Is(err, models.ErrInvalidArgument) {
			return nil, status.Error(codes.InvalidArgument, "invalid argument")
		}
		if errors.Is(err, models.ErrJsonMarshal) {
			return nil, status.Error(codes.InvalidArgument, "json marshalling error")
		}
		if errors.Is(err, models.ErrUserToSubscription) {
			return nil, status.Error(codes.InvalidArgument, "user to subscription not found")
		}
		if errors.Is(err, models.ErrSubscribe) {
			return nil, status.Error(codes.Unknown, "subscribe error")
		}

		return nil, status.Error(codes.Internal, "unknown error")
	}

	return nil, nil
}
