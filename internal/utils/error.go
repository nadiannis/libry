package utils

import "errors"

var (
	ErrBookNotFound          = errors.New("book not found")
	ErrUserNotFound          = errors.New("user not found")
	ErrUserExists            = errors.New("user already exists")
	ErrBookCurrentlyBorrowed = errors.New("book is currently borrowed")
)
