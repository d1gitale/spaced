package cli

import (
	"fmt"

	"github.com/spf13/cobra"
)

func NewDeleteCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "delete <name>",
		Short: "Delete a task",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			name := args[0]
			fmt.Printf("Deleting task: %s\n", name)
			// TODO: Implement logic
		},
	}
}
