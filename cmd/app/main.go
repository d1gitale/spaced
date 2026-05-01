package main

import (
	"log"

	"github.com/d1gitale/spaced/internal/cli"
)

func main() {
	// TODO: read envs
	// TODO: inject dependencies
	if err := cli.Execute(); err != nil {
		log.Fatal(err)
	}
}
