package models

import (
	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID   `json:"id" db:"id"`
	TgID      int64       `json:"tg_id" db:"tg_id"`
	Username  string      `json:"username" db:"username"`
	Subscribe []Subscribe `json:"subscribe" db:"subscribe"`
}

type Subscribe struct {
	TgID       int64  `json:"tg_id" db:"tg_id"`
	Username   string `json:"username" db:"username"`
	IsAccepted bool   `json:"is_accepted" db:"is_accepted"`
}

type CreateUserRequest struct {
	TgID     int64  `json:"tg_id" db:"tg_id"`
	Username string `json:"username" db:"username"`
}
