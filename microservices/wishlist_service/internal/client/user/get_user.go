package user

import (
	"context"
	"log"

	user_v1 "github.com/RSODA/wishlist/pkg/proto/user/v1"
)

func (s *userService) GetUserSub(ctx context.Context, tg_id int64) (*user_v1.GetSubResponse, error) {

	res, err := s.UserService.GetSub(ctx, &user_v1.GetSubRequest{
		TgId: tg_id,
	})

	if err != nil {
		return nil, err
	}

	log.Printf("GetSub success: tg_id=%d username=%q follower_to=%d follower_from=%d",
		tg_id,
		res.GetUsername(),
		len(res.GetFollowerTo()),
		len(res.GetFollowerFrom()),
	)

	return res, nil
}
