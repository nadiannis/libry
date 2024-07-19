package repository

import "github.com/nadiannis/libry/internal/domain"

type BookReader interface {
	GetAllBooks() []*domain.Book
	GetBookByID(bookID string) (*domain.Book, error)
}

type BookWriter interface {
	AddBook(book *domain.Book) *domain.Book
}

type IBookRepository interface {
	BookReader
	BookWriter
}
