package user

import (
	"context"
	"log"

	sqr "github.com/Masterminds/squirrel"
	"github.com/RSODA/wishlist/internal/models"
)

func (p *Postgres) GetSubIsAccepted(ctx context.Context, req *models.GetSubIsAccepted) (bool, error) {
	var res bool

	builder := sqr.Select("is_accepted").From(subTableName).Where(sqr.Eq{subFromId: req.FromTgID, subToId: req.ToTgID}).PlaceholderFormat(sqr.Dollar)
	query, args, err := builder.ToSql()
	if err != nil {
		log.Println("error building query: ", err.Error())
		return false, err
	}

	err = p.db.QueryRow(ctx, query, args...).Scan(&res)
	if err != nil {
		log.Println("err rows scan: ", err.Error())
		return false, err
	}

	return res, nil
}
