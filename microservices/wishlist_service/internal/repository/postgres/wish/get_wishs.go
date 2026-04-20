package wish

import (
	"context"
	"log"

	sqr "github.com/Masterminds/squirrel"
	errors_entity "github.com/RSODA/wishlist/microservices/wishlist_service/internal/err"
	"github.com/RSODA/wishlist/microservices/wishlist_service/internal/models"
)

func (p *postgres) GetWishs(ctx context.Context, req *models.GetWishsRequest) (*models.GetWishsResponse, error) {
	var res models.GetWishsResponse

	log.Println("request: ", req)

	builder := sqr.Select("wish.id, wish.title, wish.price, wish.url, wish.picture, status.tg_id, status.title").
		From(tableWishList).
		LeftJoin(statusTableName + " ON status.id = wish.id").
		Limit(10).
		Offset(req.Offset).
		Where(sqr.Eq{"wish.tg_id": req.TgID}).
		PlaceholderFormat(sqr.Dollar)

	query, args, err := builder.ToSql()
	if err != nil {
		log.Println("err build sql query: ", err)
		return nil, errors_entity.ErrGetWishs
	}

	rows, err := p.db.Query(ctx, query, args...)
	if err != nil {
		log.Println("err rows: ", err)
		return nil, errors_entity.ErrGetWishs
	}
	defer rows.Close()

	for rows.Next() {
		var wish models.Wish

		err = rows.Scan(&wish.ID, &wish.Title, &wish.Price, &wish.URL, &wish.Picture, &wish.Status.TgID, &wish.Status.Title)
		if err != nil {
			log.Println("err scan rows: ", err)
			return nil, errors_entity.ErrGetWishs
		}

		res.Wishs = append(res.Wishs, &wish)
	}

	err = rows.Err()
	if err != nil {
		log.Println("err rows: ", err)
		return nil, errors_entity.ErrGetWishs
	}

	log.Println("res: ", res)

	return &res, nil
}
