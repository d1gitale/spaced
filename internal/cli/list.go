package cli

import (
	"context"
	"fmt"

	"github.com/spf13/cobra"
)

func NewListCmd(ctx context.Context) *cobra.Command {
	return &cobra.Command{
		Use:   "list",
		Short: "List all tasks",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Listing all tasks...")
			// TODO: Implement logic
		},
	}
}
