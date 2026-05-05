// Package sqlite defines the SQLite adapter for retriving data
package sqlite

import (
	"context"
)

type Config struct {
	User     string `envconfig:"POSTGRES_USER"     required:"true"`
	Password string `envconfig:"POSTGRES_PASSWORD" required:"true"`
	Port     string `envconfig:"POSTGRES_PORT"     required:"true"`
	Host     string `envconfig:"POSTGRES_HOST"     required:"true"`
	DBName   string `envconfig:"POSTGRES_DB_NAME"  required:"true"`
}

type Pool struct{}

func New(ctx context.Context, c Config) (*Pool, error) {
	return &Pool{}, nil
}
