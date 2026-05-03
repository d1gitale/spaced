package cli

import (
	"context"
	"fmt"

	"github.com/spf13/cobra"
)

func NewAddCmd(ctx context.Context) *cobra.Command {
	return &cobra.Command{
		Use:   "add <name>",
		Short: "Add a new task for spaced repetition",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			name := args[0]
			fmt.Printf("Adding task: %s\n", name)
			// TODO: Implement logic
		},
	}
}
