package domain

import (
	"time"

	"github.com/google/uuid"
)

type Card struct {
	ID           uuid.UUID
	Name         string
	DueDate      time.Time
	Repetition   int
	IntervalDays int
	EaseFactor   float64
	CreatedAt    time.Time
}
