package wish

import (
	"github.com/RSODA/wishlist/microservices/wishlist_service/internal/service"
	wish_v1 "github.com/RSODA/wishlist/microservices/wishlist_service/pkg/proto/wish/v1"
)

type Implementation struct {
	wish_v1.UnimplementedWishV1Server
	wishService service.Service
}

func NewImplementation(wishService service.Service) *Implementation {
	return &Implementation{wishService: wishService}
}
