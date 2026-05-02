// Package config defines helper functions for manipulating config data
package config

import (
	"os"
	"path/filepath"
)

func DBPath() (string, error) {
	if dir := os.Getenv("XDG_DATA_HOME"); dir != "" {
		return filepath.Join("spaced"), nil
	}

	home, err := os.UserHomeDir()
	if err != nil {
		return "", nil
	}
	return filepath.Join(home, ".local", "share", "spaced", "data.db"), nil
}
