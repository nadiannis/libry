package repository

import (
	"github.com/nadiannis/libry/internal/domain"
	"github.com/nadiannis/libry/internal/utils"
)

type BookRepository struct {
	db map[string]*domain.Book
}

func NewBookRepository() IBookRepository {
	return &BookRepository{
		db: make(map[string]*domain.Book),
	}
}

func (r *BookRepository) GetAllBooks() []*domain.Book {
	books := make([]*domain.Book, 0)
	for _, book := range r.db {
		books = append(books, book)
	}
	return books
}

func (r *BookRepository) AddBook(book *domain.Book) *domain.Book {
	r.db[book.ID] = book
	return book
}

func (r *BookRepository) GetBookByID(bookID string) (*domain.Book, error) {
	book, exists := r.db[bookID]
	if !exists {
		return nil, utils.ErrBookNotFound
	}

	return book, nil
}
