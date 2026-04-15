package user

import (
	"context"
	"database/sql"
	"log"

	sqr "github.com/Masterminds/squirrel"
	customErr "github.com/RSODA/wishlist/internal/models"
	"github.com/RSODA/wishlist/internal/repository/models"
)

func (p *Postgres) GetSubFrom(ctx context.Context, tgID int64) (*models.GetSubResponse, error) {
	var items []models.Subscribe
	var username string

	err := p.db.QueryRow(
		ctx,
		"SELECT username FROM "+userTableName+" WHERE "+userTgIdColumn+" = $1",
		tgID,
	).Scan(&username)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Println("User not found: ", err)
			return nil, customErr.ErrUserNotFound
		}
		log.Println("err query username: ", err)
		return nil, customErr.ErrGetSub
	}

	builder := sqr.
		Select("sub.id_to, sub.is_accepted, u2.username AS to_username").
		From(subTableName + " sub").
		Join(userTableName + " u2 ON u2.tg_id = sub.id_to").
		Where(sqr.Eq{subFromId: tgID}).
		PlaceholderFormat(sqr.Dollar)
	query, args, err := builder.ToSql()
	if err != nil {
		log.Println("err building sql: ", err)
		return nil, customErr.ErrGetSub
	}

	rows, err := p.db.Query(ctx, query, args...)
	if err != nil {
		log.Println("error sub: ", err)
		return nil, customErr.ErrGetSub
	}
	defer rows.Close()

	for rows.Next() {
		var item models.Subscribe

		err = rows.Scan(&item.TgID, &item.IsAccepted, &item.Username)
		if err != nil {
			log.Println("err scanning rows: ", err)
			return nil, customErr.ErrGetSub
		}

		items = append(items, item)
	}

	if err = rows.Err(); err != nil {
		log.Println("rows iteration err: ", err)
		return nil, customErr.ErrGetSub
	}

	return &models.GetSubResponse{
		Username: username,
		SubFrom:  items,
	}, nil
}
