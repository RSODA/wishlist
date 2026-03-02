package api

import (
	"context"
	"errors"

	"github.com/RSODA/wishlist/internal/conventer"
	"github.com/RSODA/wishlist/internal/models"
	wishlist "github.com/RSODA/wishlist/pkg/proto/user/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (i *Implementation) CreateUser(ctx context.Context, req *wishlist.CreateUserRequest) (*wishlist.CreateUserResponse, error) {
	err := i.userService.CreateUser(ctx, conventer.ToServiceUserCreate(req))
	if err != nil {
		if errors.Is(err, models.ErrUserIsExist) {
			return nil, status.Error(codes.AlreadyExists, models.ErrUserIsExist.Error())
		}
		if errors.Is(err, models.ErrEmptyValue) {
			return nil, status.Error(codes.InvalidArgument, models.ErrEmptyValue.Error())
		}
		if errors.Is(err, models.ErrCreateUser) {
			return nil, status.Error(codes.FailedPrecondition, models.ErrCreateUser.Error())
		}

		return nil, status.Error(codes.Internal, models.ErrCreateUser.Error())
	}

	return nil, nil
}
