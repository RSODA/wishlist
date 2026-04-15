package user

import (
	"context"

	user_v1 "github.com/RSODA/wishlist/pkg/proto/user/v1"
)

type Client interface {
	GetUserSub(ctx context.Context, tg_id int64) (*user_v1.GetSubResponse, error)
}

type userService struct {
	UserService user_v1.UserV1Client
}

func NewUserService(us user_v1.UserV1Client) *userService {
	return &userService{UserService: us}
}
