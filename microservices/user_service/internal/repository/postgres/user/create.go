package user

import (
	"context"
	"log"

	sqr "github.com/Masterminds/squirrel"
	"github.com/RSODA/wishlist/internal/models"
	modelsRepo "github.com/RSODA/wishlist/internal/repository/models"
	"github.com/jackc/pgx/v5/pgconn"
)

func (p *Postgres) CreateUser(ctx context.Context, req *modelsRepo.CreateUserRequest) error {
	builder := sqr.Insert(userTableName).PlaceholderFormat(sqr.Dollar).Columns(tgIdColumn, usernameColumn).Values(req.TgID, req.Username)

	query, args, err := builder.ToSql()
	if err != nil {
		log.Println("Err convert sql: ", err)
		return models.ErrCreateUser
	}

	_, err = p.db.Exec(ctx, query, args...)
	if err != nil {
		if err.(*pgconn.PgError).Code == "23505" {
			log.Println("User already exists")
			return models.ErrUserIsExist
		}
		log.Println("Err exec: ", err)
		return models.ErrCreateUser
	}

	return nil
}
