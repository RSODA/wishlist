package repository

import (
	"context"

	"github.com/RSODA/wishlist/internal/repository/models"
)

type UserRepository interface {
	CreateUser(ctx context.Context, req *models.CreateUserRequest) error
}
