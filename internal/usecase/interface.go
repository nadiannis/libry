package usecase

import (
	"github.com/nadiannis/libry/internal/domain"
	"github.com/nadiannis/libry/internal/dto"
)

type BookReader interface {
	GetAllBooks() []*domain.Book
	GetBookByID(bookID string) (*domain.Book, error)
}

type BookWriter interface {
	AddBook(book *dto.BookInput) *domain.Book
}

type IBookUsecase interface {
	BookReader
	BookWriter
}

type UserReader interface {
	GetAllUsers() []*domain.User
	GetUserByUsername(username string) (*domain.User, error)
}

type UserWriter interface {
	AddUser(user *dto.UserInput) (*domain.User, error)
}

type IUserUsecase interface {
	UserReader
	UserWriter
}

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
