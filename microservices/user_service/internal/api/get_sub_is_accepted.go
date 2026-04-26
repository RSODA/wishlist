package api

import (
	"context"

	"github.com/RSODA/wishlist/internal/models"
	user_v1 "github.com/RSODA/wishlist/pkg/proto/user/v1"
)

func (i *Implementation) CheckSubIsAccepted(ctx context.Context, req *user_v1.CheckSubIsAcceptedRequest) (*user_v1.CheckSubIsAcceptedResponse, error) {
	res, err := i.userService.GetSubIsAccepted(ctx, &models.GetSubIsAccepted{
		FromTgID: req.FromTgId,
		ToTgID:   req.ToTgId,
	})

	if err != nil {
		return nil, err
	}

	return &user_v1.CheckSubIsAcceptedResponse{
		IsAccepted: res,
	}, nil
}
