package user

import (
	"context"
	"fmt"
	"log"

	sqr "github.com/Masterminds/squirrel"
	"github.com/RSODA/wishlist/internal/models"
	modelsRepo "github.com/RSODA/wishlist/internal/repository/models"
)

func (p *Postgres) AcceptedSub(ctx context.Context, req *modelsRepo.AcceptedSubRequests) error {
	builder := sqr.Update(subTableName).
		Set(subIsAccepted, true).
		Where(
			sqr.Eq{subToId: req.AcceptedTgId},
			sqr.Eq{subFromId: req.TgID}).
		PlaceholderFormat(sqr.Dollar)

	fmt.Println(req)
	
	query, args, err := builder.ToSql()
	if err != nil {
		log.Println("Err parse to sql: ", err)
		return models.ErrAcceptedSub
	}

	result, err := p.db.Exec(ctx, query, args...)
	if err != nil {
		log.Println("Err exec to sql: ", err)
		return models.ErrAcceptedSub
	}
	if result.RowsAffected() == 0 {
		log.Println("user not found AcceptedSub")
		return models.ErrUserNotFound
	}

	return nil
}
