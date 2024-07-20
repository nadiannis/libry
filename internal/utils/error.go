package utils

import "errors"

var (
	ErrBookNotFound          = errors.New("book not found")
	ErrBorrowedBookNotFound  = errors.New("borrowed book not found")
	ErrUserNotFound          = errors.New("user not found")
	ErrUserExists            = errors.New("user already exists")
	ErrBookCurrentlyBorrowed = errors.New("book is currently borrowed")
	ErrOverdueBookReturned   = errors.New("overdue book returned")
)
