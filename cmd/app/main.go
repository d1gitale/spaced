package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/d1gitale/spaced/internal/adapter/sqlite"
	"github.com/d1gitale/spaced/internal/cli"
	"github.com/d1gitale/spaced/pkg/logger"
)

func run(sqliteConfig sqlite.Config) {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer cancel()

	l := logger.NewLogger()

	defer func() {
		if err := l.Sync(); err != nil {
			fmt.Fprintf(os.Stderr, "failed to sync logs: %v", err)
		}
	}()

	ctx = logger.WithLogger(ctx, l)

	db, err := sqlite.New(ctx, sqliteConfig)
	if err != nil {
		l.Fatal("failed to init db: %v", err)
	}

	defer func() {
		if err := db.Close(); err != nil {
			fmt.Fprintf(os.Stderr, "failed to close db: %v", err)
		}
	}()

	rootCmd := cli.NewRootCmd(ctx, db)

	if err := rootCmd.Execute(ctx); err != nil {
		fmt.Fprintf(os.Stderr, "failed to execute: %v", err)
	}
}

func main() {
	run(sqlite.Config{
		BusyTimeout: 5000,
		JournalMode: "WAL",
		CacheSize:   8000,
		DBPath:      "data/spaced.db",
	})
}
