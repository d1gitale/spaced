package main

import (
	"context"
	"errors"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/d1gitale/spaced/internal/adapter/sqlite"
	"github.com/d1gitale/spaced/internal/cli"
	"github.com/d1gitale/spaced/pkg/config"
	"github.com/d1gitale/spaced/pkg/logger"
)

func run(sqliteConfig sqlite.Config) error {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer cancel()

	l := logger.NewLogger()

	defer func() {
		if err := l.Sync(); err != nil && !errors.Is(err, syscall.EINVAL) {
			fmt.Fprintf(os.Stderr, "failed to sync logs: %v\n", err)
		}
	}()

	ctx = logger.WithLogger(ctx, l)

	db, err := sqlite.New(ctx, sqliteConfig)
	if err != nil {
		return fmt.Errorf("failed to init db: %v", err)
	}

	defer func() {
		if err := db.Close(); err != nil {
			fmt.Fprintf(os.Stderr, "failed to close db: %v\n", err)
		}
	}()

	rootCmd := cli.NewRootCmd(ctx, db)

	if err := rootCmd.Execute(ctx); err != nil {
		return fmt.Errorf("failed to execute: %v", err)
	}

	return nil
}

func main() {
	dbPath, err := config.DBPath()
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to create directories: %v", err)
	}

	if err := run(sqlite.Config{
		BusyTimeout: 5000,
		JournalMode: "WAL",
		CacheSize:   8000,
		DBPath:      dbPath,
	}); err != nil {
		fmt.Fprintf(os.Stderr, "critical: %v\n", err)
		os.Exit(1)
	}
}
