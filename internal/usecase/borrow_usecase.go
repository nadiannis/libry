package usecase

import "github.com/nadiannis/libry/internal/domain"

type BorrowReader interface {
	GetAllBorrowedBooks() []*domain.Borrow
}

type IBorrowUsecase interface {
	BorrowReader
}
