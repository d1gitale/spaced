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

const SQLiteTimeLayout = "2006-01-02T15:04:05Z"

func New(ctx context.Context, c Config) (*Repo, error) {
	dsn := fmt.Sprintf("%s?_busy_timeout=%d&_journal_mode=%s&_foreign_keys=ON&_cache_size=%d&_time_format=datetime", c.DBPath, c.BusyTimeout, c.JournalMode, c.CacheSize)
	db, err := sql.Open("sqlite", dsn)
	if err != nil {
		l := logger.LoggerFromCtx(ctx)
		l.Fatal("failed to open DB: %v", err)
	}

	db.SetMaxOpenConns(1)
	db.SetMaxIdleConns(1)
	db.SetConnMaxLifetime(30 * time.Minute)

	if err := db.PingContext(ctx); err != nil {
		_ = db.Close()
		return nil, fmt.Errorf("failed to ping database: %w", err)
	}

	if err := runMigrations(ctx, db); err != nil {
		_ = db.Close()
		return nil, fmt.Errorf("failed to run migrations: %v", err)
	}

	return &Repo{db: db}, nil
}

func (repo *Repo) Close() error {
	return repo.db.Close()
}

func (repo *Repo) GetAllCards(ctx context.Context) ([]domain.Card, error) {
	q := `SELECT ID, name, due_date, repetition, interval_days, ease_factor FROM cards;`

	rows, err := repo.db.QueryContext(ctx, q)
	if err != nil {
		return nil, fmt.Errorf("failed to get cards from db: %v", err)
	}
	defer rows.Close()

	var cards []domain.Card
	for rows.Next() {
		card, err := scanCard(rows)
		if err != nil {
			return nil, fmt.Errorf("failed to scan card: %v", err)
		}
		cards = append(cards, card)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("failed while iterating over results: %v", err)
	}

	return cards, nil
}

func (repo *Repo) GetDueCards(ctx context.Context) ([]domain.Card, error) {
	panic("not implemented") // TODO: Implement
}

func (repo *Repo) CreateCard(ctx context.Context, r domain.Card) error {
	q := `INSERT INTO cards (ID, name, due_date, repetition, interval_days, ease_factor) VALUES (?, ?, ?, ?, ?, ?);`

	rows, err := repo.db.QueryContext(
		ctx, q,
		r.ID,
		r.Name,
		r.DueDate.Format(SQLiteTimeLayout),
		r.Repetition,
		r.IntervalDays,
		r.EaseFactor,
	)
	if err != nil {
		return fmt.Errorf("failed to insert card into db: %v", err)
	}
	defer rows.Close()

	return nil
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
	CREATE TABLE IF NOT EXISTS cards (
    ID             UUID PRIMARY KEY,
    name           TEXT NOT NULL,
    due_date       DATE NOT NULL,
    repetition     INTEGER NOT NULL DEFAULT 0 CHECK (repetition >= 0),
    interval_days  INTEGER NOT NULL DEFAULT 0 CHECK (interval_days >= 0),
    ease_factor    REAL    NOT NULL DEFAULT 2.5 CHECK (ease_factor >= 1.3)
	);
	CREATE INDEX IF NOT EXISTS idx_cards_due ON cards (due_date ASC);
	`

	_, err := db.ExecContext(ctx, schema)
	return err
}

func scanCard(rows *sql.Rows) (domain.Card, error) {
	var c domain.Card
	var idStr, dueDateStr string

	err := rows.Scan(
		&idStr,
		&c.Name,
		&dueDateStr,
		&c.Repetition,
		&c.IntervalDays,
		&c.EaseFactor,
	)
	if err != nil {
		return domain.Card{}, fmt.Errorf("failed to scan card row: %w", err)
	}

	c.ID, err = uuid.Parse(idStr)
	if err != nil {
		return domain.Card{}, fmt.Errorf("failed to parse uuid '%s': %w", idStr, err)
	}

	c.DueDate, err = time.Parse(SQLiteTimeLayout, dueDateStr)
	if err != nil {
		c.DueDate = time.Time{}
	}

	return c, nil
}
