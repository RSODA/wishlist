package repository

import (
	"context"

	"github.com/RSODA/wishlist/internal/models"
	repoModels "github.com/RSODA/wishlist/internal/repository/models"
)

type UserRepository interface {
	CreateUser(ctx context.Context, req *repoModels.CreateUserRequest) error
	GetSub(ctx context.Context, tgID int64) (*[]models.Subscribe, error)
}
