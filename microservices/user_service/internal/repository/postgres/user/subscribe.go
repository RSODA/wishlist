package user

import (
	"context"
	"log"

	sqr "github.com/Masterminds/squirrel"
	"github.com/RSODA/wishlist/internal/models"
	modelsRepository "github.com/RSODA/wishlist/internal/repository/models"
)

func (p *Postgres) Subscribe(ctx context.Context, req *modelsRepository.SubscribeRequest) error {
	var exist int

	query := `
	SELECT 1
WHERE NOT EXISTS (
    SELECT 1
    FROM ` + subTableName + `
    WHERE id_from = $1 AND id_to = $2
)`

	err := p.db.QueryRow(ctx, query, req.TgID, req.ToTgID).Scan(&exist)

	log.Println(err, req.ToTgID)

	if err != nil {
		if err.Error() == "no rows in result set" {
			log.Println("User already subscribed: ", req.TgID, req.ToTgID)
			return models.ErrUserAlreadySubscribed
		}
		return err
	}

	builder := sqr.Insert(subTableName).PlaceholderFormat(sqr.Dollar).Values(req.TgID, req.ToTgID).Columns(subFromId, subToId)
	query, args, err := builder.ToSql()
	if err != nil {
		log.Println("err build query for Subscribe: ", err)
		return models.ErrSubscribe
	}

	_, err = p.db.Exec(ctx, query, args...)
	if err != nil {
		log.Printf("failed to execute query: %v", err)
		return models.ErrSubscribe
	}

	return nil
}
