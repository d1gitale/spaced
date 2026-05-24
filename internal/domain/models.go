package domain

import (
	"github.com/google/uuid"
)

type Card struct {
	ID           uuid.UUID
	Name         string
	DueDate      string
	Repetition   int
	IntervalDays int
	EaseFactor   float64
}
