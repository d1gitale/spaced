// Package cli handles commands parsing and calls business logic
package cli

import "github.com/spf13/cobra"

// NewRootCmd creates the root command and adds all subcommands.
func NewRootCmd() *cobra.Command {
	rootCmd := &cobra.Command{
		Use:   "spaced",
		Short: "Spaced repetition CLI",
		Long:  "spaced is a spaced repetition tracker that helps organizing your schedule and works well with Waybar",
	}

	rootCmd.AddCommand(
		NewAddCmd(),
		NewListCmd(),
		NewCheckCmd(),
		NewRenameCmd(),
		NewDeleteCmd(),
	)

	return rootCmd
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() error {
	return NewRootCmd().Execute()
}
