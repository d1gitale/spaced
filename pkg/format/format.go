// Package format implements functions for printing output in different formats
package format

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/d1gitale/spaced/internal/domain"
)

func PrintCard(ctx context.Context, card *domain.Card, flag string) error {
	switch flag {
	case "json":
		err := printJSON(card)
		if err != nil {
			return fmt.Errorf("failed to print JSON: %v", err)
		}
	default:
		fmt.Println(*card)
	}

	return nil
}

func printJSON(card *domain.Card) error {
	json, err := json.Marshal(card)
	if err != nil {
		return fmt.Errorf("failed to marshal into JSON: %v", err)
	}

	fmt.Println(string(json))

	return nil
}
