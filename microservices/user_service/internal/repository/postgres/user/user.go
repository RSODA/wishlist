package user

import (
	"github.com/RSODA/wishlist/internal/repository"
	"github.com/jackc/pgx/v5/pgxpool"
)

const (
	userTableName      = "users"
	userTgIdColumn     = "tg_id"
	userUsernameColumn = "username"

	subTableName  = "sub"
	subFromId     = "id_from"
	subToId       = "id_to"
	subIsAccepted = "is_accepted"
)

type Postgres struct {
	db *pgxpool.Pool
}

func NewPostgres(db *pgxpool.Pool) repository.UserRepository {
	return &Postgres{db: db}
}
