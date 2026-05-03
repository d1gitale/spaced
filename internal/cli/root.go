// Package cli handles commands parsing and calls business logic
package cli

import (
	"context"

	"github.com/spf13/cobra"
)

// NewRootCmd creates the root command and adds all subcommands.
func NewRootCmd(ctx context.Context) *cobra.Command {
	rootCmd := &cobra.Command{
		Use:   "github.com/d1gitale/spaced",
		Short: "Spaced repetition CLI",
		Long:  "spaced is a spaced repetition tracker that helps organizing your schedule and works well with Waybar",
	}

	rootCmd.AddCommand(
		NewAddCmd(ctx),
		NewListCmd(ctx),
		NewCheckCmd(ctx),
		NewRenameCmd(ctx),
		NewDeleteCmd(ctx),
	)

	return rootCmd
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute(ctx context.Context) error {
	return NewRootCmd(ctx).Execute()
}
