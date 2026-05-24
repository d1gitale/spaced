package cli

import (
	"fmt"
	"time"

	"github.com/d1gitale/spaced/internal/domain"
	"github.com/d1gitale/spaced/pkg/constants"
	"github.com/d1gitale/spaced/pkg/sm2"
	"github.com/spf13/cobra"
)

func NewAddCmd(repo domain.CardAdapter) *cobra.Command {
	return &cobra.Command{
		Use:   "add <name>",
		Short: "Add a new task for spaced repetition",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx := cmd.Context()

			name := args[0]
			repetition := 0
			ef := 2.5
			interval := sm2.GetInterval(ef, repetition, 0)
			due := time.Now().AddDate(0, 0, interval).Format(constants.SpacedDateFmt)

			card := domain.Card{
				Name:         name,
				Repetition:   repetition,
				EaseFactor:   ef,
				IntervalDays: interval,
				DueDate:      due,
			}

			err := repo.CreateCard(ctx, card)
			if err != nil {
				return fmt.Errorf("failed to create card: %v", err)
			}

			return nil
		},
	}
}
