package models

type Wish struct {
	ID      int64  `json:"id" db:"id"`
	TgID    int64  `json:"tg_id" db:"tg_id"`
	Title   string `json:"title" db:"title"`
	Price   int64  `json:"price" db:"price"`
	URL     string `json:"url" db:"url"`
	Picture string `json:"picture" db:"picture"`
	Status  Status `json:"status" db:"status"`
}

type Status struct {
	ID    int64   `json:"id" db:"id"`
	Title *string `json:"title" db:"title"`
	TgID  *int64  `json:"tg_id" db:"tg_id"`
}

type CreateRequest struct {
	TgID    int64  `json:"tg_id" validate:"required"`
	Title   string `json:"title" validate:"required"`
	URL     string `json:"url" validate:"url"`
	Price   int64  `json:"price" validate:"required"`
	Picture string `json:"picture" validate:"datauri"`
}

type GetWishsRequest struct {
	TgID   int64  `json:"tg_id" validate:"required"`
	ToTgID int64  `json:"to_tg_id" validate:"required"`
	Offset uint64 `json:"offset" validate:"required"`
}

type GetWishsResponse struct {
	Wishs []*Wish `json:"wishes"`
}
