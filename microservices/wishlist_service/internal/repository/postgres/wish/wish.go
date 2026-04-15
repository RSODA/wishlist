package wish

import (
	"github.com/RSODA/wishlist/microservices/wishlist_service/internal/repository"
	"github.com/jackc/pgx/v5/pgxpool"
)

const (
	tableWishList             = "wish"
	idColumnWishList          = "id"
	tgIdColumnWishList        = "tg_id"
	titleColumnWishList       = "title"
	descriptionColumnWishList = "description"
	priceColumnWishList       = "price"
	urlColumnWishList         = "url"
	pictureColumnWishList     = "picture"

	statusTableName    = "status"
	wishIdColumnStatus = "id"
	titleColumnStatus  = "title"
	tgIdColumnStatus   = "tg_id"
)

type postgres struct {
	db *pgxpool.Pool
}

func New(db *pgxpool.Pool) (repository.Repository, error) {
	return &postgres{
		db: db,
	}, nil
}
