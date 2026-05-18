// Package sqlite defines the SQLite adapter for retriving data
package sqlite

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/d1gitale/spaced/internal/domain"
	"github.com/d1gitale/spaced/pkg/logger"
	"github.com/google/uuid"
	_ "modernc.org/sqlite"
)

type Config struct {
	BusyTimeout uint   `envconfig:"SQLITE_BUSY_TIMEOUT" reuired:"true"`
	JournalMode string `envconfig:"SQLITE_JOURNAL_MODE" required:"true"`
	CacheSize   int    `envconfig:"SQLITE_CACHE_SIZE" required:"true"`
	DBPath      string `envconfig:"SQLITE_DB_PATH"  required:"true"`
}

type Repo struct {
	db *sql.DB
}

func New(ctx context.Context, c Config) (*Repo, error) {
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

	return &Repo{db: db}, nil
}

func (repo *Repo) Close() error {
	return repo.db.Close()
}

func (repo *Repo) GetAllCards(ctx context.Context) ([]domain.Card, error) {
	panic("not implemented") // TODO: Implement
}

func (repo *Repo) GetDueCards(ctx context.Context) ([]domain.Card, error) {
	panic("not implemented") // TODO: Implement
}

func (repo *Repo) CreateCard(ctx context.Context, r domain.Card) error {
	panic("not implemented") // TODO: Implement
}

func (repo *Repo) MarkReviewed(ctx context.Context, id uuid.UUID, easinessFactor float64, interval int, repetition int, due time.Time) error {
	panic("not implemented") // TODO: Implement
}

func (repo *Repo) RenameCard(ctx context.Context, id uuid.UUID, newName string) error {
	panic("not implemented") // TODO: Implement
}

func (repo *Repo) RemoveCard(ctx context.Context, id uuid.UUID) error {
	panic("not implemented") // TODO: Implement
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
