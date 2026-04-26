package wish

import (
	"context"
	"log"

	sqr "github.com/Masterminds/squirrel"
	"github.com/RSODA/wishlist/microservices/wishlist_service/internal/err"
	"github.com/RSODA/wishlist/microservices/wishlist_service/internal/models"
)

func (p *postgres) Create(ctx context.Context, req *models.CreateRequest) (*int64, error) {
	var id int64

	builder := sqr.Insert(tableWishList).Columns(tgIdColumnWishList, titleColumnWishList, priceColumnWishList, urlColumnWishList, pictureColumnWishList).
		PlaceholderFormat(sqr.Dollar).
		Values(req.TgID, req.Title, req.Price, req.URL, req.Picture).Suffix("RETURNING id")

	query, args, err := builder.ToSql()
	if err != nil {
		log.Println("err build sql Create Wish query: ", err)
		return nil, errors_entity.ErrCreateWish
	}

	err = p.db.QueryRow(ctx, query, args...).Scan(&id)
	if err != nil {
		log.Println("err query Wish List: ", err)
	}

	return &id, nil
}
