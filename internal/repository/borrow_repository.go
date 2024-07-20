package repository

import (
	"time"

	"github.com/nadiannis/libry/internal/domain"
)

type BorrowReader interface {
	GetAllBorrowedBooks() []*domain.Borrow
	GetBorrowedBook(userID, bookID string) (*domain.Borrow, error)
}

type BorrowWriter interface {
	AddBorrowedBook(borrow *domain.Borrow) (*domain.Borrow, error)
	UpdateStatus(borrowID, status string) (*domain.Borrow, error)
	UpdateDates(borrowID string, startDate, endDate time.Time) error
}

type IBorrowRepository interface {
	BorrowReader
	BorrowWriter
}
