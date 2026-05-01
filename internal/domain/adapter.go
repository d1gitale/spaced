// Package domain declares domain types
package domain

type ReviewDB interface {
	GetAll()
	GetDue()
	AddReview(name string)
	DeleteReview(id int)
	RenameReview(id int, new string)
	CheckReview(id int)
}
