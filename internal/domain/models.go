package domain

type Card struct {
	ID           int64   `json:"id"`
	Name         string  `json:"name"`
	DueDate      string  `json:"dueDate"`
	Repetition   int     `json:"repetition"`
	IntervalDays int     `json:"interval"`
	EaseFactor   float64 `json:"EF"`
}
