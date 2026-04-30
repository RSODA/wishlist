package config

import (
	"errors"
	"os"
)

type PostgresConfig interface {
	DSN() string
}
type postgresConfig struct {
	dsn string
}

func NewPostgresConfig() (PostgresConfig, error) {
	dsn := os.Getenv("PG_DSN")
	if len(dsn) == 0 {
		return nil, errors.New("PG_DSN environment variable not set")
	}
	return &postgresConfig{dsn: dsn}, nil
}

func (pc *postgresConfig) DSN() string {
	return pc.dsn
}
