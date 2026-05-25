package cli

import (
	"fmt"

	"github.com/d1gitale/spaced/internal/domain"
	"github.com/spf13/cobra"
)

func NewRenameCmd(repo domain.CardAdapter) *cobra.Command {
	return &cobra.Command{
		Use:   "rename --id <id> --name <name>",
		Short: "Rename a task",
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx := cmd.Context()

			id, err := cmd.Flags().GetInt64("id")
			if err != nil {
				fmt.Println("invalid id")
				return fmt.Errorf("failed to get --id flag value: %v", err)
			}
			name, err := cmd.Flags().GetString("name")
			if err != nil {
				fmt.Println("invalid name")
				return fmt.Errorf("failed to get --name flag value: %v", err)
			}

			err = repo.RenameCard(ctx, id, name)
			if err != nil {
				return fmt.Errorf("failed to rename card: %v", err)
			}

			return nil
		},
	}
}
