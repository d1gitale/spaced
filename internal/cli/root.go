// Package cli handles commands parsing and calls business logic
package cli

import (
	"context"
	"fmt"

	"github.com/d1gitale/spaced/internal/domain"
	"github.com/spf13/cobra"
)

type RootCmd struct {
	root *cobra.Command
	repo domain.CardAdapter
}

func NewRootCmd(ctx context.Context, repo domain.CardAdapter) (*RootCmd, error) {
	rootCmd := &RootCmd{}

	rootCmd.repo = repo

	rootCmd.root = &cobra.Command{
		Use:   "spaced",
		Short: "Spaced repetition CLI",
		Long:  "spaced is a spaced repetition tracker that helps organizing your schedule and works well with Waybar",
	}

	listCmd := NewListCmd(repo)
	listCmd.Flags().String("format", "", "format to print cards in")

	checkCmd := NewCheckCmd(repo)
	checkCmd.Flags().Int64P("id", "i", 0, "integer id of a card to check")
	err := checkCmd.MarkFlagRequired("id")
	if err != nil {
		return nil, fmt.Errorf("looks like there is a typo: %v", err)
	}

	checkCmd.Flags().IntP("score", "p", 5, "integer retention score for checked card")
	err = checkCmd.MarkFlagRequired("score")
	if err != nil {
		return nil, fmt.Errorf("looks like there is a typo: %v", err)
	}

	rootCmd.root.AddCommand(
		NewAddCmd(repo),
		listCmd,
		checkCmd,
		NewRenameCmd(repo),
		NewDeleteCmd(repo),
	)

	return rootCmd, nil
}

func (rootCmd *RootCmd) Execute(ctx context.Context) error {
	return rootCmd.root.ExecuteContext(ctx)
}
