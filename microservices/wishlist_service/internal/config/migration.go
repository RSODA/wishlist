package config

import (
	"log"
	"os"
)

type MigrationsConfig interface {
	MigrationPath() string
}

type migrationsConfig struct {
	migrationPath string
}

func NewMigrationsConfig() MigrationsConfig {
	path := os.Getenv("MIGRATIONS_PATH")
	if len(path) == 0 {
		log.Fatalf("environment variable MIGRATIONS_PATH not defined")
		return nil
	}

	return &migrationsConfig{
		migrationPath: path,
	}
}

func (m *migrationsConfig) MigrationPath() string {
	return m.migrationPath
}
