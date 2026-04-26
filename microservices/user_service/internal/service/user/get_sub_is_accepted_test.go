package user

import (
	"context"
	"errors"
	"testing"

	"github.com/RSODA/wishlist/internal/models"
	repoModels "github.com/RSODA/wishlist/internal/repository/models"
	"github.com/jackc/pgx/v5"
)

type userRepositoryStub struct {
	getSubIsAcceptedFn func(ctx context.Context, req *models.GetSubIsAccepted) (bool, error)
}

func (r *userRepositoryStub) CreateUser(context.Context, *repoModels.CreateUserRequest) error {
	panic("unexpected CreateUser call")
}

func (r *userRepositoryStub) GetSubFrom(context.Context, int64) (*repoModels.GetSubResponse, error) {
	panic("unexpected GetSubFrom call")
}

func (r *userRepositoryStub) GetSubTo(context.Context, int64) ([]repoModels.Subscribe, error) {
	panic("unexpected GetSubTo call")
}

func (r *userRepositoryStub) Subscribe(context.Context, *repoModels.SubscribeRequest) error {
	panic("unexpected Subscribe call")
}

func (r *userRepositoryStub) AcceptedSub(context.Context, *repoModels.AcceptedSubRequests) error {
	panic("unexpected AcceptedSub call")
}

func (r *userRepositoryStub) GetUser(context.Context, string) (int64, error) {
	panic("unexpected GetUser call")
}

func (r *userRepositoryStub) GetSubIsAccepted(ctx context.Context, req *models.GetSubIsAccepted) (bool, error) {
	if r.getSubIsAcceptedFn == nil {
		return false, nil
	}

	return r.getSubIsAcceptedFn(ctx, req)
}

func TestGetSubIsAccepted_ReturnsFalseWhenSubscriptionRowDoesNotExist(t *testing.T) {
	service := &userService{
		repo: &userRepositoryStub{
			getSubIsAcceptedFn: func(context.Context, *models.GetSubIsAccepted) (bool, error) {
				return false, pgx.ErrNoRows
			},
		},
	}

	res, err := service.GetSubIsAccepted(context.Background(), &models.GetSubIsAccepted{
		FromTgID: 100,
		ToTgID:   200,
	})
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if res {
		t.Fatal("expected subscription to be treated as not accepted")
	}
}

func TestGetSubIsAccepted_ReturnsRepositoryValue(t *testing.T) {
	service := &userService{
		repo: &userRepositoryStub{
			getSubIsAcceptedFn: func(context.Context, *models.GetSubIsAccepted) (bool, error) {
				return true, nil
			},
		},
	}

	res, err := service.GetSubIsAccepted(context.Background(), &models.GetSubIsAccepted{
		FromTgID: 100,
		ToTgID:   200,
	})
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if !res {
		t.Fatal("expected subscription to be accepted")
	}
}

func TestGetSubIsAccepted_ReturnsUnexpectedRepositoryErrors(t *testing.T) {
	expectedErr := errors.New("db unavailable")
	service := &userService{
		repo: &userRepositoryStub{
			getSubIsAcceptedFn: func(context.Context, *models.GetSubIsAccepted) (bool, error) {
				return false, expectedErr
			},
		},
	}

	_, err := service.GetSubIsAccepted(context.Background(), &models.GetSubIsAccepted{
		FromTgID: 100,
		ToTgID:   200,
	})
	if !errors.Is(err, expectedErr) {
		t.Fatalf("expected %v, got %v", expectedErr, err)
	}
}
