package user

import (
	"context"

	"github.com/RSODA/wishlist/internal/models"
	"github.com/RSODA/wishlist/internal/repository"
)

type UserService interface {
	CreateUser(ctx context.Context, req *models.CreateUserRequest) error
	GetSub(ctx context.Context, tgID int64) (*[]models.Subscribe, error)
	Subscribe(ctx context.Context, req *models.SubscribeRequest) error
	AcceptedSub(ctx context.Context, req *models.AcceptedSubRequest) error
}

type userService struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) UserService {
	return &userService{
		repo: repo,
	}
}
