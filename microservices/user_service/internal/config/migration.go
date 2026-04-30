package config

import (
	"errors"
	"os"
)

type MigrationsConfig interface {
	MigrationPath() string
}

type migrationsConfig struct {
	migrationPath string
}

func NewMigrationsConfig() (MigrationsConfig, error) {
	path := os.Getenv("MIGRATIONS_PATH")
	if len(path) == 0 {
		return nil, errors.New("MIGRATIONS_PATH environment variable not set")
	}

	return &migrationsConfig{
		migrationPath: path,
	}, nil
}

func (m *migrationsConfig) MigrationPath() string {
	return m.migrationPath
}
