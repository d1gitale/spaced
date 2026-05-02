// Package adapter implements adapters for data extraction via Adapter interface from the domain package
package adapter

import (
	"context"
	"database/sql"
	"fmt"
	"os"
	"path/filepath"

	"github.com/d1gitale/spaced/internal/config"
)

type Pool struct{}

func New() (*sql.DB, func(), error) {
	path, err := config.DBPath()
	if err != nil {
		return nil, nil, fmt.Errorf("failed to get DB path: %v", err)
	}
	if err := os.MkdirAll(filepath.Dir(path), os.ModeDir); err != nil {
		return nil, nil, fmt.Errorf("failed to create DB directory: %v", err)
	}

	dsn := fmt.Sprintf("file:%s?_busy_timeout=3000&_journal_mode=WAL&_foreign_keys=ON", path)
	pool, err := sql.Open("sqlite", dsn)
	if err != nil {
		return nil, nil, err
	}
	pool.SetMaxOpenConns(1)

	cleanup := func() { pool.Close() }
	if err := pool.PingContext(context.Background()); err != nil {
		return nil, nil, fmt.Errorf("failed to ping DB: %v", err)
	}

	return pool, cleanup, nil
}
