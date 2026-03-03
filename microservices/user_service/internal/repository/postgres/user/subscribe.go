package user

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"log"

	sqr "github.com/Masterminds/squirrel"
	"github.com/RSODA/wishlist/internal/models"
	modelsRepository "github.com/RSODA/wishlist/internal/repository/models"
)

func (p *Postgres) Subscribe(ctx context.Context, req *modelsRepository.SubscribeRequest) error {
	sub := models.Subscribe{
		TgID:       req.ToTgID,
		Username:   req.ToUsername,
		IsAccepted: false,
	}

	data, err := json.Marshal([]models.Subscribe{sub})
	if err != nil {
		log.Printf("failed to marshal subscribe request: %v", err)
		return models.ErrJsonMarshal
	}

	builder := sqr.Update(userTableName).PlaceholderFormat(sqr.Dollar).Set(subscribeColumn,
		sqr.Expr("COALESCE("+subscribeColumn+", '[]'::jsonb) || ?::jsonb", data),
	).Where(sqr.Eq{tgIdColumn: req.TgID})

	query, args, err := builder.ToSql()
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return models.ErrUserNotFound
		}
		log.Printf("failed to build query: %v", err)
		return models.ErrSubscribe
	}

	_, err = p.db.Exec(ctx, query, args...)
	if err != nil {
		log.Printf("failed to execute query: %v", err)
		return models.ErrSubscribe
	}

	return nil
}
