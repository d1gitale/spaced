// Package sqlite defines the SQLite adapter for retriving data
package sqlite

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/d1gitale/spaced/pkg/logger"
	_ "github.com/mattn/go-sqlite3"
)

type Config struct {
	BusyTimeout uint   `envconfig:"SQLITE_BUSY_TIMEOUT" reuired:"true"`
	JournalMode string `envconfig:"SQLITE_JOURNAL_MODE" required:"true"`
	CacheSize   int    `envconfig:"SQLITE_CACHE_SIZE" required:"true"`
	DBPath      string `envconfig:"SQLITE_DB_PATH"  required:"true"`
}

type Pool struct {
	db *sql.DB
}

func New(ctx context.Context, c Config) (*Pool, error) {
	dsn := fmt.Sprintf("%s?_busy_timeout=%d&_journal_mode=%s&_foreign_keys=ON&_cache_size=%d", c.DBPath, c.BusyTimeout, c.JournalMode, c.CacheSize)
	db, err := sql.Open("go-sqlite3", dsn)
	if err != nil {
		l := logger.LoggerFromCtx(ctx)
		l.Fatal("failed to open DB: %v", err)
	}

	db.SetMaxOpenConns(1)
	db.SetMaxIdleConns(1)
	db.SetConnMaxLifetime(30 * time.Minute)

	if err := runMigrations(ctx, db); err != nil {
		_ = db.Close()
		l := logger.LoggerFromCtx(ctx)
		l.Fatal("failed to run migrations: %v", err)
	}

	return &Pool{db: db}, nil
}

func runMigrations(ctx context.Context, db *sql.DB) error {
	return nil
}
