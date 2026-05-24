// Package domain defines domain-specific entities
package domain

import (
	"context"
)

type CardAdapter interface {
	GetAllCards(ctx context.Context) ([]Card, error)
	GetCardByID(ctx context.Context, id int64) (Card, error)
	GetDueCards(ctx context.Context) ([]Card, error)
	CreateCard(ctx context.Context, r Card) error
	MarkReviewed(ctx context.Context, id int64, easinessFactor float64, interval int, repetition int, due string) error
	RenameCard(ctx context.Context, id int64, newName string) error
	RemoveCard(ctx context.Context, id int64) error
}
