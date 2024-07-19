package repository

import (
	"github.com/nadiannis/libry/internal/domain"
	"github.com/nadiannis/libry/internal/utils"
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

func (r *BorrowRepository) AddBorrowedBook(borrow *domain.Borrow) (*domain.Borrow, error) {
	for _, borrowedBook := range r.db {
		if borrow.BookID == borrowedBook.BookID && utils.TimeIsBetween(borrow.StartDate, borrowedBook.StartDate, borrowedBook.EndDate) {
			return nil, utils.ErrBookCurrentlyBorrowed
		}
	}

	r.db[borrow.ID] = borrow
	return borrow, nil
}
