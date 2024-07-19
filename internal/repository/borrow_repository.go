package repository

import "github.com/nadiannis/libry/internal/domain"

type BorrowReader interface {
	GetAllBorrowedBooks() []*domain.Borrow
}

type BorrowWriter interface {
	AddBorrowedBook(borrow *domain.Borrow) (*domain.Borrow, error)
}

type IBorrowRepository interface {
	BorrowReader
	BorrowWriter
}
