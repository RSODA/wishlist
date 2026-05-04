package user

import (
	"context"
	"log"

	sqr "github.com/Masterminds/squirrel"
	customErr "github.com/RSODA/wishlist/internal/models"
	"github.com/RSODA/wishlist/internal/repository/models"
)

func (p *Postgres) GetSubTo(ctx context.Context, tgId int64) ([]models.Subscribe, error) {
	var res []models.Subscribe

	builder := sqr.Select("sub.id_from, sub.is_accepted, users.username AS from_username").
		From(userTableName + " users").
		PlaceholderFormat(sqr.Dollar).
		Join(subTableName + " sub ON users.tg_id = sub.id_from").
		Where(sqr.Eq{"sub.id_to": tgId})

	query, args, err := builder.ToSql()
	if err != nil {
		log.Println("error building query getsubto", err)
		return nil, customErr.ErrGetSub
	}

	rows, err := p.db.Query(ctx, query, args...)
	if err != nil {
		log.Println("error executing query getsubto", err)
		return nil, customErr.ErrGetSub
	}

	for rows.Next() {
		var item models.Subscribe

		err = rows.Scan(&item.TgID, &item.IsAccepted, &item.Username)
		if err != nil {
			log.Println("error scanning row", err)
			return nil, customErr.ErrGetSub
		}

		res = append(res, item)
	}

	defer rows.Close()

	log.Println("result get sub to: ", res)

	return res, nil
}
