package user

import (
	"context"
	"database/sql"
	"errors"
	"log"

	sqr "github.com/Masterminds/squirrel"
	"github.com/RSODA/wishlist/internal/models"
	modelsRepo "github.com/RSODA/wishlist/internal/repository/models"
)

func (p *Postgres) AcceptedSub(ctx context.Context, req *modelsRepo.AcceptedSubRequests) error {
	builder := sqr.Update(userTableName).Where(sqr.Eq{tgIdColumn: req.TgID}).Set(subscribeColumn, sqr.Expr(`
		(
			SELECT jsonb_agg(
				CASE
					WHEN (elem->>'tg_id')::bigint = ?
					THEN jsonb_set(elem, '{is_accepted}', 'true'::jsonb)
					ELSE elem
				END
			)
			FROM jsonb_array_elements(subscribe) AS elem
		)
	`, req.AcceptedTgId)).PlaceholderFormat(sqr.Dollar)

	query, args, err := builder.ToSql()
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			log.Println("user not found AcceptedSub: ", err)
			return models.ErrUserNotFound
		}
		log.Println("Err parse to sql: ", err)
		return models.ErrAcceptedSub
	}

	_, err = p.db.Exec(ctx, query, args...)
	if err != nil {
		log.Println("Err exec to sql: ", err)
		return models.ErrAcceptedSub
	}

	return nil
}
