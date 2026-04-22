package cli

import (
	"fmt"

	"github.com/spf13/cobra"
)

func NewRenameCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "rename <old_name> <new_name>",
		Short: "Rename a task",
		Args:  cobra.ExactArgs(2),
		Run: func(cmd *cobra.Command, args []string) {
			oldName := args[0]
			newName := args[1]
			fmt.Printf("Renaming task from %s to %s\n", oldName, newName)
			// TODO: Implement logic
		},
	}
}
