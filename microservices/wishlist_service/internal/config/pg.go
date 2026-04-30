package config

import (
	"errors"
	"os"
)

type PGConfig interface {
	DSN() string
}

type pgConfig struct {
	Dsn string
}

func NewPGConfig() (PGConfig, error) {
	dsn := os.Getenv("PG_DSN")
	if len(dsn) == 0 {
		return nil, errors.New("PG_DSN environment variable is required")
	}

	return &pgConfig{
		Dsn: dsn,
	}, nil
}

func (p *pgConfig) DSN() string {
	return p.Dsn
}
