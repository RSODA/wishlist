package wish

import (
	"context"
	"log"

	sqr "github.com/Masterminds/squirrel"
	"github.com/RSODA/wishlist/microservices/wishlist_service/internal/err"
	"github.com/RSODA/wishlist/microservices/wishlist_service/internal/models"
)

func (p *postgres) GetById(ctx context.Context, id int64) (*models.Wish, error) {
	var wish models.Wish

	builder := sqr.Select("wish.id, wish.tg_id, wish.title, wish.price, wish.url, wish.picture, status.id, status.title, status.tg_id").
		From("wish").
		Where(sqr.Eq{"wish.id": id}).
		LeftJoin(statusTableName + " ON status.id = wish.id")

	query, args, err := builder.PlaceholderFormat(sqr.Dollar).ToSql()
	if err != nil {
		log.Println("err building sql get wish by id: ", err)
		return nil, errors_entity.ErrGetWishById
	}

	err = p.db.QueryRow(ctx, query, args...).Scan(&wish.ID, &wish.TgID, &wish.Title, &wish.Price, &wish.URL, &wish.Picture, &wish.Status.ID, &wish.Title, &wish.TgID)
	if err != nil {
		log.Println("err executing sql get wish by id: ", err)
	}

	return &wish, nil
}
