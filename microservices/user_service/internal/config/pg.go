package config

import (
	"log"
	"os"
)

type PostgresConfig interface {
	DSN() string
}
type postgresConfig struct {
	dsn string
}

func NewPostgresConfig() PostgresConfig {
	dsn := os.Getenv("PG_DSN")
	if len(dsn) == 0 {
		log.Fatal("environment variable dsn is not set")
		return nil
	}
	return &postgresConfig{dsn: dsn}
}

func (pc *postgresConfig) DSN() string {
	return pc.dsn
}
