package domain

import "time"

const (
	StatusBorrowed = "borrowed"
	StatusReturned = "returned"
	StatusOverdue  = "overdue"
)

type BorrowStatus struct {
	Status string
}

type Borrow struct {
	ID        string
	BookID    string
	UserID    string
	StartDate time.Time
	EndDate   time.Time
	BorrowStatus
}
