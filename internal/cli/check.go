package cli

import (
	"fmt"

	"github.com/d1gitale/spaced/internal/domain"
	"github.com/spf13/cobra"
)

func NewCheckCmd(repo domain.CardAdapter) *cobra.Command {
	return &cobra.Command{
		Use:   "check <name> <grade>",
		Short: "Update retention grade for a task",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			name := args[0]
			grade := args[1]
			fmt.Printf("Checking task %s with grade %s\n", name, grade)
			return nil
			// TODO: Implement logic
		},
	}
}
