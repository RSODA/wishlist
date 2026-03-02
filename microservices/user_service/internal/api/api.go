package api

import (
	"github.com/RSODA/wishlist/internal/service/user"
	wishlist "github.com/RSODA/wishlist/pkg/proto/user/v1"
)

type Implementation struct {
	wishlist.UnimplementedUserV1Server
	userService user.UserService
}

func NewImplementation(userService user.UserService) *Implementation {
	return &Implementation{
		userService: userService,
	}
}
