package user

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"

	sqr "github.com/Masterminds/squirrel"
	customErr "github.com/RSODA/wishlist/internal/models"
	"github.com/RSODA/wishlist/internal/repository/models"
)

func (p *Postgres) GetSub(ctx context.Context, tgID int64, isAccepted bool) (*models.GetSubResponse, error) {
	var items []models.Subscribe
	var username string

	fmt.Println("hi: ", tgID, isAccepted)

	builder := sqr.Select("users.username, sub.id_to, sub.is_accepted, u2.username AS to_username").From(userTableName + " users").PlaceholderFormat(sqr.Dollar).LeftJoin(subTableName + " sub ON users.tg_id = sub.id_from").
		LeftJoin(userTableName + " u2 ON u2.tg_id = sub.id_to").
		Where(sqr.Eq{
			"users." + userTgIdColumn: tgID,
			"sub." + subIsAccepted:    isAccepted,
		})
	query, args, err := builder.ToSql()
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			log.Println("User not found: ", err)
			return nil, customErr.ErrUserNotFound
		}
		log.Println("err building sql: ", err)
		return nil, customErr.ErrGetSub
	}

	rows, err := p.db.Query(ctx, query, args...)
	if err != nil {
		fmt.Println("error sub: ", err)
		return nil, customErr.ErrGetSub
	}

	for rows.Next() {
		var item models.Subscribe

		err = rows.Scan(&username, &item.TgID, &item.IsAccepted, &item.Username)
		if err != nil {
			log.Println("err scanning rows: ", err)
			return nil, customErr.ErrGetSub
		}

		items = append(items, item)
	}

	return &models.GetSubResponse{
		Username: username,
		Sub:      items,
	}, nil
}
