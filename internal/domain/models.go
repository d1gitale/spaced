package domain

import (
	"time"

	"github.com/google/uuid"
)

type Card struct {
	ID             uuid.UUID
	Name           string
	Due            time.Time
	Repetition     int
	Interval       int
	EasinessFactor float64
}
