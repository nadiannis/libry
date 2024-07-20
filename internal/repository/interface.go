package repository

import (
	"time"

	"github.com/nadiannis/libry/internal/domain"
)

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

type UserReader interface {
	GetAllUsers() []*domain.User
	GetUserByUsername(username string) (*domain.User, error)
}

type UserWriter interface {
	AddUser(user *domain.User) *domain.User
	AddBook(userID string, book *domain.Book) (*domain.Book, error)
	DeleteBookByID(userID, bookID string) error
}

type IUserRepository interface {
	UserReader
	UserWriter
}

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
