// Package sqlite defines the SQLite adapter for retriving data
package sqlite

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/d1gitale/spaced/pkg/logger"
	_ "modernc.org/sqlite"
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

func (db *Pool) Close() error {
	return db.db.Close()
}

func runMigrations(ctx context.Context, db *sql.DB) error {
	schema := `
	CREATE TABLE cards (
    id             UUID PRIMARY KEY, -- или INTEGER AUTOINCREMENT для SQLite
    name           TEXT NOT NULL,
    due_date       DATE NOT NULL DEFAULT CURRENT_DATE,
    repetition     INTEGER NOT NULL DEFAULT 0 CHECK (repetition >= 0),
    interval_days  INTEGER NOT NULL DEFAULT 0 CHECK (interval_days >= 0),
    ease_factor    REAL    NOT NULL DEFAULT 2.5 CHECK (ease_factor >= 1.3),
    created_at     TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	);
	CREATE INDEX idx_cards_due ON cards (due_date ASC);
	`

	_, err := db.ExecContext(ctx, schema)
	return err
}
