package repository

import (
	"context"

	"github.com/RSODA/wishlist/microservices/wishlist_service/internal/models"
)

type Repository interface {
	Create(ctx context.Context, req *models.CreateRequest) (*int64, error)
	GetById(ctx context.Context, id int64) (*models.Wish, error)
}
