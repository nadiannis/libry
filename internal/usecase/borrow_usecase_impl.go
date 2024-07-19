package usecase

import (
	"github.com/nadiannis/libry/internal/domain"
	"github.com/nadiannis/libry/internal/repository"
)

type BorrowUsecase struct {
	repository repository.IBorrowRepository
}

func NewBorrowUsecase(repository repository.IBorrowRepository) IBorrowUsecase {
	return &BorrowUsecase{
		repository: repository,
	}
}

func (u *BorrowUsecase) GetAllBorrowedBooks() []*domain.Borrow {
	return u.repository.GetAllBorrowedBooks()
}
