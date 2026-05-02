// Package domain declares domain types
package domain

type ReviewDB interface {
	GetAll() []Review
	GetDue() []Review
	AddReview(name string)
	DeleteReview(id int)
	RenameReview(id int, new string)
	CheckReview(id int)
}
