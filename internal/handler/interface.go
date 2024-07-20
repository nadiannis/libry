package handler

import "github.com/nadiannis/libry/internal/dto"

type BookReader interface {
	GetAllBooks(parts []string)
}

type BookWriter interface {
	AddBook(input *dto.BookInput)
}

type IBookHandler interface {
	BookReader
	BookWriter
}

type UserReader interface {
	GetAllUsers(parts []string)
}

type UserWriter interface {
	AddUser(input *dto.UserInput)
}

type IUserHandler interface {
	UserReader
	UserWriter
}

type BorrowReader interface {
	GetAllBorrowedBooks(parts []string)
}

type BorrowWriter interface {
	BorrowBook(parts []string)
	ReturnBook(parts []string)
}

type IBorrowHandler interface {
	BorrowReader
	BorrowWriter
}
