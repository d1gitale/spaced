package domain

import (
	"github.com/google/uuid"
)

type Card struct {
	ID           uuid.UUID `json:"id"`
	Name         string    `json:"name"`
	DueDate      string    `json:"dueDate"`
	Repetition   int       `json:"repetition"`
	IntervalDays int       `json:"interval"`
	EaseFactor   float64   `json:"EF"`
}
