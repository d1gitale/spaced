// Package domain defines domain-specific entities
package domain

import (
	"context"
	"time"

	"github.com/google/uuid"
)

type CardAdapter interface {
	GetAllCards(ctx context.Context) ([]Card, error)
	GetDueCards(ctx context.Context) ([]Card, error)
	CreateCard(ctx context.Context, r Card) error
	MarkReviewed(ctx context.Context, id uuid.UUID, easinessFactor float64, interval int, repetition int, due time.Time) error
	RenameCard(ctx context.Context, id uuid.UUID, newName string) error
	RemoveCard(ctx context.Context, id uuid.UUID) error
}
