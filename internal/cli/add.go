package cli

import (
	"fmt"

	"github.com/d1gitale/spaced/internal/domain"
	"github.com/spf13/cobra"
)

func NewAddCmd(repo domain.CardAdapter) *cobra.Command {
	return &cobra.Command{
		Use:   "add <name>",
		Short: "Add a new task for spaced repetition",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			name := args[0]
			fmt.Printf("Adding task: %s\n", name)
			// TODO: Implement logic
			return nil
		},
	}
}
