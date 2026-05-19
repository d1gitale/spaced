// Package cli handles commands parsing and calls business logic
package cli

import (
	"context"

	"github.com/d1gitale/spaced/internal/domain"
	"github.com/spf13/cobra"
)

type RootCmd struct {
	root *cobra.Command
	repo domain.CardAdapter
}

func NewRootCmd(ctx context.Context, repo domain.CardAdapter) *RootCmd {
	rootCmd := &RootCmd{}

	rootCmd.repo = repo

	rootCmd.root = &cobra.Command{
		Use:   "github.com/d1gitale/spaced",
		Short: "Spaced repetition CLI",
		Long:  "spaced is a spaced repetition tracker that helps organizing your schedule and works well with Waybar",
	}

	rootCmd.root.AddCommand(
		NewAddCmd(repo),
		NewListCmd(repo),
		NewCheckCmd(repo),
		NewRenameCmd(repo),
		NewDeleteCmd(repo),
	)

	return rootCmd
}

func (rootCmd *RootCmd) Execute(ctx context.Context) error {
	return rootCmd.root.ExecuteContext(ctx)
}
