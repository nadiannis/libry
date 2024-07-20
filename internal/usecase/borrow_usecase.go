package usecase

import (
	"github.com/nadiannis/libry/internal/domain"
	"github.com/nadiannis/libry/internal/dto"
)

type BorrowReader interface {
	GetAllBorrowedBooks() []*domain.Borrow
}

type BorrowWriter interface {
	BorrowBook(input *dto.BorrowInput) (*domain.Borrow, error)
	ReturnBook(input *dto.BorrowInput) (*domain.Borrow, error)
}

type IBorrowUsecase interface {
	BorrowReader
	BorrowWriter
}
