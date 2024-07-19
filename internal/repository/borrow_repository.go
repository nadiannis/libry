package repository

import "github.com/nadiannis/libry/internal/domain"

type BorrowReader interface {
	GetAllBorrowedBooks() []*domain.Borrow
}

type IBorrowRepository interface {
	BorrowReader
}
