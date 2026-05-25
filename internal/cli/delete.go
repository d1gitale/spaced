package cli

import (
	"fmt"

	"github.com/d1gitale/spaced/internal/domain"
	"github.com/spf13/cobra"
)

func NewDeleteCmd(repo domain.CardAdapter) *cobra.Command {
	return &cobra.Command{
		Use:   "delete --id <id>",
		Short: "Delete a task",
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx := cmd.Context()
			id, err := cmd.Flags().GetInt64("id")
			if err != nil {
				fmt.Println("invalid id")
				return fmt.Errorf("failed to get --id flag: %v", err)
			}

			err = repo.RemoveCard(ctx, id)
			if err != nil {
				return fmt.Errorf("failed to remove card %d: %v", id, err)
			}

			return nil
		},
	}
}
