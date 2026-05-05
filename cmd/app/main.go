package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/d1gitale/spaced/internal/cli"
	"github.com/d1gitale/spaced/pkg/logger"
)

func main() {
	l := logger.NewLogger()
	defer func() {
		if err := l.Sync(); err != nil {
			fmt.Fprintf(os.Stderr, "failed to sync logs: %v", err)
		}
	}()

	ctx := logger.WithLogger(context.Background(), l)

	if err := cli.Execute(ctx); err != nil {
		log.Fatal(err)
	}
}
