package cli

import (
	"fmt"

	"github.com/d1gitale/spaced/internal/domain"
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

			// TODO: move into pkg func with formatting
			for i, c := range cards {
				fmt.Printf("card %d: %v", i, c)
			}
			return nil
		},
	}
}
