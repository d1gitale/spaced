package cli

import (
	"fmt"

	"github.com/d1gitale/spaced/internal/domain"
	"github.com/spf13/cobra"
)

func NewRenameCmd(repo domain.CardAdapter) *cobra.Command {
	return &cobra.Command{
		Use:   "rename <old_name> <new_name>",
		Short: "Rename a task",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			oldName := args[0]
			newName := args[1]
			fmt.Printf("Renaming task from %s to %s\n", oldName, newName)
			return nil
			// TODO: Implement logic
		},
	}
}
