package user

import (
	"context"
	"log"
	"strings"

	"github.com/RSODA/wishlist/internal/conventer"
	serviceModels "github.com/RSODA/wishlist/internal/models"
)

func (r *userService) CreateUser(ctx context.Context, req *serviceModels.CreateUserRequest) error {
	if len(strings.TrimSpace(req.Username)) == 0 || req.TgID < 0 {
		log.Printf("invalid username: %v or tg_id: %v", req.Username, req.TgID)
		return serviceModels.ErrEmptyValue
	}

	err := r.repo.CreateUser(ctx, conventer.ToRepoUserCreate(req))
	if err != nil {
		return err
	}

	return nil
}
