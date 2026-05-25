// Package sqlite defines the SQLite adapter for retriving data
package sqlite

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/d1gitale/spaced/internal/domain"
	"github.com/d1gitale/spaced/pkg/constants"
	"github.com/d1gitale/spaced/pkg/logger"
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

func (repo *Repo) GetCardByID(ctx context.Context, id int64) (domain.Card, error) {
	q := `SELECT ID, name, due_date, repetition, interval_days, ease_factor FROM cards WHERE ID = ?`

	row := repo.db.QueryRowContext(ctx, q, id)

	card := domain.Card{}
	var dueDateStr string
	err := row.Scan(
		&card.ID,
		&card.Name,
		&dueDateStr,
		&card.Repetition,
		&card.IntervalDays,
		&card.EaseFactor,
	)
	if err != nil {
		return domain.Card{}, fmt.Errorf("failed to get card %d from db: %v", id, err)
	}

	parsedDue, err := time.Parse(constants.SQLiteTimeLayout, dueDateStr)
	if err != nil {
		return domain.Card{}, fmt.Errorf("failed to parse due_date from db: %v", err)
	}

	card.DueDate = parsedDue.Format(constants.SpacedDateFmt)

	return card, nil
}

func (repo *Repo) GetDueCards(ctx context.Context) ([]domain.Card, error) {
	cards, err := repo.GetAllCards(ctx)
	if err != nil {
		return nil, err
	}

	var resSet []domain.Card
	for _, c := range cards {
		due, err := time.Parse(constants.SpacedDateFmt, c.DueDate)
		if err != nil {
			return nil, fmt.Errorf("failed to parse c.DueDate: %v", err)
		}

		year, month, day := time.Now().Date()
		isDueToday := due.Local().Compare(time.Date(year, month, day, 0, 0, 0, 0, time.UTC).Local())

		if isDueToday != 1 {
			resSet = append(resSet, c)
		}
	}

	return resSet, nil
}

func (repo *Repo) CreateCard(ctx context.Context, r domain.Card) error {
	q := `INSERT INTO cards (name, due_date, repetition, interval_days, ease_factor) VALUES (?, ?, ?, ?, ?);`

	stmt, err := repo.db.PrepareContext(ctx, q)
	if err != nil {
		return fmt.Errorf("failed to prepare statement: %v", err)
	}

	dueDate, err := time.Parse(constants.SpacedDateFmt, r.DueDate)
	if err != nil {
		return fmt.Errorf("failed to parse date from model: %v", err)
	}

	_, err = stmt.ExecContext(
		ctx,
		r.Name,
		dueDate.Format(constants.SQLiteTimeLayout),
		r.Repetition,
		r.IntervalDays,
		r.EaseFactor,
	)
	if err != nil {
		return fmt.Errorf("failed to insert card into db: %v", err)
	}

	return nil
}

func (repo *Repo) MarkReviewed(ctx context.Context, id int64, easinessFactor float64, interval int, repetition int, due string) error {
	q := "UPDATE cards SET due_date = ?, repetition = ?, interval_days = ?, ease_factor = ? WHERE ID = ?;"

	stmt, err := repo.db.PrepareContext(ctx, q)
	if err != nil {
		return fmt.Errorf("failed to prepare statement: %v", err)
	}

	_, err = stmt.ExecContext(
		ctx,
		due,
		repetition,
		interval,
		easinessFactor,
		id,
	)
	if err != nil {
		return fmt.Errorf("failed to mark card reviewed in db: %v", err)
	}

	return nil
}

func (repo *Repo) RenameCard(ctx context.Context, id int64, newName string) error {
	q := "UPDATE cards SET name = ? WHERE ID = ?;"

	stmt, err := repo.db.PrepareContext(ctx, q)
	if err != nil {
		return fmt.Errorf("failed to prepare statement: %v", err)
	}

	_, err = stmt.ExecContext(ctx, newName, id)
	if err != nil {
		return fmt.Errorf("failed to rename card in db: %v", err)
	}

	return nil
}

func (repo *Repo) RemoveCard(ctx context.Context, id int64) error {
	q := "DELETE FROM cards WHERE ID = ?;"

	stmt, err := repo.db.PrepareContext(ctx, q)
	if err != nil {
		return fmt.Errorf("failed to prepare statement: %v", err)
	}

	_, err = stmt.ExecContext(ctx, id)
	if err != nil {
		return fmt.Errorf("failed to remove card from db: %v", err)
	}

	return nil
}

func runMigrations(ctx context.Context, db *sql.DB) error {
	schema := `
	CREATE TABLE IF NOT EXISTS cards (
    ID             INTEGER PRIMARY KEY AUTOINCREMENT,
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
	var dueDateStr string

	err := rows.Scan(
		&c.ID,
		&c.Name,
		&dueDateStr,
		&c.Repetition,
		&c.IntervalDays,
		&c.EaseFactor,
	)
	if err != nil {
		return domain.Card{}, fmt.Errorf("failed to scan card row: %w", err)
	}

	fmtDate, err := time.Parse(constants.SQLiteTimeLayout, dueDateStr)
	if err != nil {
		return domain.Card{}, fmt.Errorf("failed to parse date from db: %v", err)
	}
	c.DueDate = fmtDate.Format(constants.SpacedDateFmt)

	return c, nil
}
