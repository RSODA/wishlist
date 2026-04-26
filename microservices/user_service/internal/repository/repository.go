package repository

import (
	"context"

	"github.com/RSODA/wishlist/internal/models"
	repoModels "github.com/RSODA/wishlist/internal/repository/models"
)

type UserRepository interface {
	CreateUser(ctx context.Context, req *repoModels.CreateUserRequest) error
	GetSubFrom(ctx context.Context, tgID int64) (*repoModels.GetSubResponse, error)
	GetSubTo(ctx context.Context, tgId int64) ([]repoModels.Subscribe, error)
	Subscribe(ctx context.Context, req *repoModels.SubscribeRequest) error
	AcceptedSub(ctx context.Context, req *repoModels.AcceptedSubRequests) error
	GetUser(ctx context.Context, name string) (int64, error)
	GetSubIsAccepted(ctx context.Context, req *models.GetSubIsAccepted) (bool, error)
}
