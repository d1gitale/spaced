// Package config defines helper functions for config creation
package config

import (
	"fmt"
	"os"
	"path/filepath"
)

func DBPath() (string, error) {
	dataDir := os.Getenv("XDG_DATA_HOME")
	if dataDir == "" {
		dataDir = filepath.Join(os.Getenv("HOME"), ".local", "share")
	}
	appDataDir := filepath.Join(dataDir, "spaced-rep")

	if err := os.MkdirAll(appDataDir, 0o755); err != nil {
		return "", fmt.Errorf("create data dir: %v", err)
	}

	path := filepath.Join(appDataDir, "spaced.db")
	return path, nil
}
