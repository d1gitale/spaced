package cli

import (
	"fmt"
	"time"

	"github.com/d1gitale/spaced/internal/domain"
	"github.com/d1gitale/spaced/pkg/constants"
	"github.com/d1gitale/spaced/pkg/sm2"
	"github.com/spf13/cobra"
)

func NewCheckCmd(repo domain.CardAdapter) *cobra.Command {
	return &cobra.Command{
		Use:   "check --id <id> --score <1-5>",
		Short: "Update retention score for a task",
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx := cmd.Context()
			id, err := cmd.Flags().GetInt64("id")
			if err != nil {
				fmt.Println("invalid id")
				return fmt.Errorf("failed to get --id flag value: %v", err)
			}

			score, err := cmd.Flags().GetInt("score")
			if err != nil {
				fmt.Println("invalid score")
				return fmt.Errorf("failed to get --score flag value: %v", err)
			}

			if score > 5 || score < 1 {
				fmt.Println("score needs to be an integer between 1 and 5 inclusively")
				return fmt.Errorf("score is out of bounds")
			}

			card, err := repo.GetCardByID(ctx, id)
			if err != nil {
				return fmt.Errorf("failed to get card %d by id: %v", id, err)
			}

			dueParsed, err := time.Parse(constants.SpacedDateFmt, card.DueDate)
			if err != nil {
				return fmt.Errorf("failed to parse card.DueDate: %v", err)
			}

			year, month, day := time.Now().Date()
			isDueToday := dueParsed.Local().Compare(time.Date(year, month, day, 0, 0, 0, 0, time.UTC).Local())
			if isDueToday != 1 {
				newDue := dueParsed.AddDate(0, 0, card.IntervalDays).Format(constants.SpacedDateFmt)

				newInterval := sm2.GetInterval(card.EaseFactor, card.Repetition, card.IntervalDays)
				newEF := sm2.GetEF(card.EaseFactor, score)
				newRepetition := card.Repetition + 1

				err = repo.MarkReviewed(ctx, id, newEF, newInterval, newRepetition, newDue)
				if err != nil {
					return fmt.Errorf("failed to mark card %d reviewed: %v", id, err)
				}
			} else {
				return fmt.Errorf("card is not due today")
			}

			return nil
		},
	}
}
