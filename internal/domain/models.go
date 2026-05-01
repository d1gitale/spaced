package domain

import "time"

type Review struct {
	ID         int64
	Name       string
	Due        time.Time
	Repetition int
	Interval   int
	EF         float64
}
