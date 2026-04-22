package cli

import (
	"fmt"

	"github.com/spf13/cobra"
)

func NewCheckCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "check <name> <grade>",
		Short: "Update retention grade for a task",
		Args:  cobra.ExactArgs(2),
		Run: func(cmd *cobra.Command, args []string) {
			name := args[0]
			grade := args[1]
			fmt.Printf("Checking task %s with grade %s\n", name, grade)
			// TODO: Implement logic
		},
	}
}
