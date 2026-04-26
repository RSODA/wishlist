package config

import (
	"os"

	errors_entity "github.com/RSODA/wishlist/microservices/wishlist_service/internal/err"
)

type HTTPConfig interface {
	Address() string
}

type httpConfig struct {
	Host string
	Port string
}

func NewHTTPConfig() (HTTPConfig, error) {
	host := os.Getenv("HTTP_HOST")
	if len(host) == 0 {
		return nil, errors_entity.ErrHostNotFound
	}

	port := os.Getenv("HTTP_PORT")
	if len(port) == 0 {
		return nil, errors_entity.ErrPortNotFound
	}

	return &httpConfig{
		Host: host,
		Port: port,
	}, nil
}

func (c *httpConfig) Address() string {
	return c.Host + ":" + c.Port
}
