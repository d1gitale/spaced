package main

import (
	"context"
	"log"

	"github.com/d1gitale/spaced/internal/cli"
	"github.com/d1gitale/spaced/pkg/logger"
)

func main() {
	l := logger.NewLogger()
	defer func() {
		if err := l.Sync(); err != nil {
			log.Fatalf("failed to sync logs: %v", err)
		}
	}()

	ctx := logger.WithLogger(context.Background(), l)

	if err := cli.Execute(ctx); err != nil {
		log.Fatal(err)
	}
}
