package user

import (
	"context"
	"testing"

	"github.com/RSODA/wishlist/internal/conventer"
	"github.com/RSODA/wishlist/internal/models"
	"github.com/RSODA/wishlist/internal/repository"
	"github.com/RSODA/wishlist/internal/repository/mocks"
	repoModels "github.com/RSODA/wishlist/internal/repository/models"
	"github.com/gojuno/minimock/v3"
	"github.com/stretchr/testify/require"
)

func TestSubscribe(t *testing.T) {
	t.Parallel()
	type userRepoMockFunc func(mc *minimock.Controller, req *repoModels.SubscribeRequest) repository.UserRepository

	type args struct {
		ctx context.Context
		req *models.SubscribeRequest
	}

	ctx := context.Background()

	tests := []struct {
		name         string
		args         args
		err          error
		repoMockFunc userRepoMockFunc
	}{
		{
			name: "success",
			args: args{
				ctx: ctx,
				req: &models.SubscribeRequest{
					TgID:       1,
					ToTgID:     2,
					ToUsername: "some_username",
				},
			},
			err: nil,
			repoMockFunc: func(mc *minimock.Controller, req *repoModels.SubscribeRequest) repository.UserRepository {
				mock := mocks.NewUserRepositoryMock(mc)
				mock.GetUserMock.Expect(ctx, "some_username").Return(2, nil)
				mock.SubscribeMock.Expect(ctx, req).Return(nil)
				return mock
			},
		},
		{
			name: "invalid tg_id",
			args: args{
				ctx: ctx,
				req: &models.SubscribeRequest{
					TgID:   -1,
					ToTgID: 2,
				},
			},
			err: models.ErrInvalidArgument,
		},
		{
			name: "identicalID",
			args: args{
				ctx: ctx,
				req: &models.SubscribeRequest{
					TgID:   1,
					ToTgID: 1,
				},
			},
			err: models.ErrIdenticalID,
		},
		{
			name: "get user error",
			args: args{
				ctx: ctx,
				req: &models.SubscribeRequest{
					TgID:       1,
					ToTgID:     0,
					ToUsername: "some_user",
				},
			},
			err: models.ErrGetUser, // или какую ошибку возвращает твой repo
			repoMockFunc: func(mc *minimock.Controller, req *repoModels.SubscribeRequest) repository.UserRepository {
				mock := mocks.NewUserRepositoryMock(mc)
				mock.GetUserMock.Expect(ctx, "some_user").Return(0, models.ErrGetUser)
				return mock
			},
		},
		{
			name: "subscribe repo error",
			args: args{
				ctx: ctx,
				req: &models.SubscribeRequest{
					TgID:       1,
					ToTgID:     0,
					ToUsername: "some_user",
				},
			},
			err: models.ErrSubscribe,
			repoMockFunc: func(mc *minimock.Controller, req *repoModels.SubscribeRequest) repository.UserRepository {
				mock := mocks.NewUserRepositoryMock(mc)
				mock.GetUserMock.Expect(ctx, "some_user").Return(0, nil)
				mock.SubscribeMock.Expect(ctx, req).Return(models.ErrSubscribe)
				return mock
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			mc := minimock.NewController(t)

			var repo repository.UserRepository
			if tt.repoMockFunc != nil {
				repoReq := conventer.ToRepoSubscribe(tt.args.req)
				repo = tt.repoMockFunc(mc, repoReq)
			}

			svc := NewUserService(repo)
			err := svc.Subscribe(tt.args.ctx, tt.args.req)
			require.Equal(t, tt.err, err)
		})
	}
}
