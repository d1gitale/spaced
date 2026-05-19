package cli

import (
	"fmt"

	"github.com/d1gitale/spaced/internal/domain"
	"github.com/spf13/cobra"
)

func NewDeleteCmd(repo domain.CardAdapter) *cobra.Command {
	return &cobra.Command{
		Use:   "delete <name>",
		Short: "Delete a task",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			name := args[0]
			fmt.Printf("Deleting task: %s\n", name)
			return nil
			// TODO: Implement logic
		},
	}
}
