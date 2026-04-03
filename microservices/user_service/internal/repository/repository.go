package repository

import (
	"context"

	repoModels "github.com/RSODA/wishlist/internal/repository/models"
)

type UserRepository interface {
	CreateUser(ctx context.Context, req *repoModels.CreateUserRequest) error
	GetSub(ctx context.Context, tgID int64, isAccepted bool) (*repoModels.GetSubResponse, error)
	Subscribe(ctx context.Context, req *repoModels.SubscribeRequest) error
	AcceptedSub(ctx context.Context, req *repoModels.AcceptedSubRequests) error
	GetUser(ctx context.Context, name string) (int64, error)
}
