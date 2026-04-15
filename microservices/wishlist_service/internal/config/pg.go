package config

import (
	"log"
	"os"
)

type PGConfig interface {
	DSN() string
}

type pgConfig struct {
	Dsn string
}

func NewPGConfig() PGConfig {
	dsn := os.Getenv("PG_DSN")
	if len(dsn) == 0 {
		log.Fatalf("env PG_DSN required")
	}

	return &pgConfig{
		Dsn: dsn,
	}
}

func (p *pgConfig) DSN() string {
	return p.Dsn
}
