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
	Username   string `json:"username" db:"to_username"`
	IsAccepted bool   `json:"is_accepted" db:"is_accepted"`
}

type GetSubResponse struct {
	Username string      `json:"username" db:"username"`
	Sub      []Subscribe `json:"sub" db:"sub"`
}
type CreateUserRequest struct {
	TgID     int64  `json:"tg_id" db:"tg_id"`
	Username string `json:"username" db:"username"`
}

type SubscribeRequest struct {
	TgID   int64 `json:"tg_id" db:"tg_id"`
	ToTgID int64 `json:"to_tg_id" db:"to_tg_id"`
}

type AcceptedSubRequests struct {
	TgID         int64 `json:"tg_id" db:"tg_id"`
	AcceptedTgId int64 `json:"accepted_tg_id" db:"accepted_tg_id"`
}
