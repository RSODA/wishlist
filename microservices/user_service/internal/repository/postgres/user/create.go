package user

import (
	"context"
	"log"

	sqr "github.com/Masterminds/squirrel"
	models2 "github.com/RSODA/wishlist/internal/models"
	"github.com/RSODA/wishlist/internal/repository/models"
	"github.com/jackc/pgx/v5/pgconn"
)

func (p *Postgres) CreateUser(ctx context.Context, req *models.CreateUserRequest) error {
	builder := sqr.Insert(userTableName).PlaceholderFormat(sqr.Dollar).Columns(tgIdColumn, usernameColumn).Values(req.TgID, req.Username)

	query, args, err := builder.ToSql()
	if err != nil {
		log.Println("Err convert sql: ", err)
		return models2.ErrCreateUser
	}

	_, err = p.db.Exec(ctx, query, args...)
	if err != nil {
		if err.(*pgconn.PgError).Code == "23505" {
			log.Println("User already exists")
			return models2.ErrUserIsExist
		}
		log.Println("Err exec: ", err)
		return models2.ErrCreateUser
	}

	return nil
}
