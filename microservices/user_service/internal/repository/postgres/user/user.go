package user

import (
	"github.com/RSODA/wishlist/internal/repository"
	"github.com/jackc/pgx/v5/pgxpool"
)

const (
	userTableName   = "users"
	idColumn        = "id"
	tgIdColumn      = "tg_id"
	usernameColumn  = "username"
	subscribeColumn = "subscribe"
)

type Postgres struct {
	db *pgxpool.Pool
}

func NewPostgres(db *pgxpool.Pool) repository.UserRepository {
	return &Postgres{db: db}
}
