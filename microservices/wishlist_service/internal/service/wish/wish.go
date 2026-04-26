package wish

import (
	"github.com/RSODA/wishlist/microservices/wishlist_service/internal/client/user"
	"github.com/RSODA/wishlist/microservices/wishlist_service/internal/repository"
	"github.com/RSODA/wishlist/microservices/wishlist_service/internal/service"
	"github.com/go-playground/validator/v10"
)

type wish struct {
	r        repository.Repository
	client   user.Client
	validate *validator.Validate
	host     string
}

func NewWishService(r repository.Repository, c user.Client, host string) service.Service {
	return &wish{r: r, client: c, validate: validator.New(), host: host}
}
