package conventer

import (
	"github.com/RSODA/wishlist/microservices/wishlist_service/internal/models"
	wish_v1 "github.com/RSODA/wishlist/microservices/wishlist_service/pkg/proto/wish/v1"
)

func ToAPIFromService(request *models.GetWishsResponse) *wish_v1.GetWishsResponse {
	var resp wish_v1.GetWishsResponse

	for _, v := range request.Wishs {
		if v == nil {
			continue
		}

		var wish wish_v1.Wish

		wish.Id = v.ID
		wish.Title = v.Title
		wish.Price = v.Price
		wish.Url = v.URL
		wish.Picture = v.Picture
		wish.TgId = v.TgID
		wish.Status = &wish_v1.Status{}

		if v.Status.TgID != nil {
			wish.Status.TgId = *v.Status.TgID
		}
		if v.Status.Title != nil {
			wish.Status.Status = *v.Status.Title
		}

		resp.Wishs = append(resp.Wishs, &wish)
	}

	return &resp
}
