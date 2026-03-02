package user

import (
	"context"
	"database/sql"
	"errors"
	"log"

	sqr "github.com/Masterminds/squirrel"
	"github.com/RSODA/wishlist/internal/models"
)

func (p *Postgres) GetSub(ctx context.Context, tgID int64) (*[]models.Subscribe, error) {
	var items []models.Subscribe

	builder := sqr.Select(subscribeColumn).From(userTableName).PlaceholderFormat(sqr.Dollar).Where(sqr.Eq{tgIdColumn: tgID})
	query, args, err := builder.ToSql()
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			log.Println("User not found: ", err)
			return nil, models.ErrUserNotFound
		}
		log.Println("err building sql: ", err)
		return nil, models.ErrGetSub
	}

	err = p.db.QueryRow(ctx, query, args...).Scan(&items)
	if err != nil {
		log.Println("err scanning rows: ", err)
		return nil, models.ErrGetSub
	}

	return &items, nil
}
