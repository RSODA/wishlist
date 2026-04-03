package user

import (
	"context"
	"errors"
	"log"

	sqr "github.com/Masterminds/squirrel"
	"github.com/RSODA/wishlist/internal/models"
	"github.com/jackc/pgx/v5"
)

func (p *Postgres) GetUser(ctx context.Context, name string) (int64, error) {
	var id int64

	builder := sqr.Select(userTgIdColumn).
		From(userTableName).
		Where(sqr.Eq{userUsernameColumn: name}).
		Limit(1).
		PlaceholderFormat(sqr.Dollar)

	query, args, err := builder.ToSql()
	if err != nil {
		log.Println("err building query: ", err.Error())
		return -1, models.ErrSubscribe
	}

	err = p.db.QueryRow(ctx, query, args...).Scan(&id)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			log.Println("user to subscription not found")
			return -1, models.ErrUserToSubscription
		}

		log.Println("failed to find user by username: ", err.Error())
		return -1, models.ErrSubscribe
	}

	return id, nil
}
