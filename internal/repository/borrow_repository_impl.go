package repository

import (
	"github.com/nadiannis/libry/internal/domain"
)

type BorrowRepository struct {
	db map[string]*domain.Borrow
}

func NewBorrowRepository() IBorrowRepository {
	return &BorrowRepository{
		db: make(map[string]*domain.Borrow),
	}
}

func (r *BorrowRepository) GetAllBorrowedBooks() []*domain.Borrow {
	borrowedBooks := make([]*domain.Borrow, 0)
	for _, borrowedBook := range r.db {
		borrowedBooks = append(borrowedBooks, borrowedBook)
	}
	return borrowedBooks
}
