package domain

import "time"

type Borrow struct {
	ID        string
	BookID    string
	UserID    string
	StartDate time.Time
	EndDate   time.Time
}
