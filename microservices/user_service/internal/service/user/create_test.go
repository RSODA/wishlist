package user

import (
	"context"
	"testing"

	serviceModels "github.com/RSODA/wishlist/internal/models"
	"github.com/RSODA/wishlist/internal/repository"
	"github.com/RSODA/wishlist/internal/repository/mocks"
	repoModels "github.com/RSODA/wishlist/internal/repository/models"
	"github.com/gojuno/minimock/v3"
	"github.com/stretchr/testify/require"
)

func TestCreate(t *testing.T) {
	t.Parallel()

	type userRepoMockFunc func(mc *minimock.Controller, req *repoModels.CreateUserRequest) repository.UserRepository

	type args struct {
		ctx context.Context
		req *serviceModels.CreateUserRequest
	}

	ctx := context.Background()

	tests := []struct {
		name     string
		args     args
		err      error
		repoMock userRepoMockFunc
	}{
		{
			name: "success",
			args: args{
				ctx: ctx,
				req: &serviceModels.CreateUserRequest{
					TgID:     1,
					Username: "username",
				},
			},
			err: nil,
			repoMock: func(mc *minimock.Controller, req *repoModels.CreateUserRequest) repository.UserRepository {
				mock := mocks.NewUserRepositoryMock(mc)
				mock.CreateUserMock.Expect(ctx, req).Return(nil)
				return mock
			},
		},
		{
			name: "fail username is null",
			args: args{
				ctx: ctx,
				req: &serviceModels.CreateUserRequest{
					TgID:     1,
					Username: "",
				},
			},
			err: serviceModels.ErrEmptyValue,
		},
		{
			name: "fail tg_id is empty",
			args: args{
				ctx: ctx,
				req: &serviceModels.CreateUserRequest{
					TgID:     -1,
					Username: "username",
				},
			},
			err: serviceModels.ErrEmptyValue,
		},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			mc := minimock.NewController(t)
			t.Cleanup(mc.Finish)

			var userRepo repository.UserRepository
			if tt.repoMock != nil {
				userRepo = tt.repoMock(mc, &repoModels.CreateUserRequest{
					TgID:     tt.args.req.TgID,
					Username: tt.args.req.Username,
				})
			}

			service := NewUserService(userRepo)

			err := service.CreateUser(tt.args.ctx, tt.args.req)
			require.Equal(t, tt.err, err)
		})
	}
}
