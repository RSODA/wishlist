package user

import (
	"context"

	"github.com/RSODA/wishlist/internal/models"
	"github.com/RSODA/wishlist/internal/repository"
	repoModels "github.com/RSODA/wishlist/internal/repository/models"
)

type UserService interface {
	CreateUser(ctx context.Context, req *models.CreateUserRequest) error
	GetSub(ctx context.Context, tgID int64) (*repoModels.GetSubResponse, error)
	Subscribe(ctx context.Context, req *models.SubscribeRequest) error
	AcceptedSub(ctx context.Context, req *models.AcceptedSubRequest) error
	GetSubIsAccepted(ctx context.Context, req *models.GetSubIsAccepted) (bool, error)
}

type userService struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) UserService {
	return &userService{
		repo: repo,
	}
}
