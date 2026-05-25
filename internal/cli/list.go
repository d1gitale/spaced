package cli

import (
	"fmt"

	"github.com/d1gitale/spaced/internal/domain"
	"github.com/d1gitale/spaced/pkg/format"
	"github.com/spf13/cobra"
)

func NewListCmd(repo domain.CardAdapter) *cobra.Command {
	return &cobra.Command{
		Use:   "list",
		Short: "List all tasks",
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx := cmd.Context()
			cards, err := repo.GetAllCards(ctx)
			if err != nil {
				return fmt.Errorf("failed to list cards: %v", err)
			}

			fmtFlag, err := cmd.Flags().GetString("format")
			if err != nil {
				fmt.Println("invalid format value")
				return fmt.Errorf("failed to get --format flag value: %v", err)
			}

			for _, c := range cards {
				err := format.PrintCard(ctx, &c, fmtFlag)
				if err != nil {
					return fmt.Errorf("failed to print card: %v", err)
				}
			}

			return nil
		},
	}
}
