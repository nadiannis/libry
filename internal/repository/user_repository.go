package repository

import "github.com/nadiannis/libry/internal/domain"

type UserReader interface {
	GetAllUsers() []*domain.User
	GetUserByUsername(username string) (*domain.User, error)
}

type UserWriter interface {
	AddUser(user *domain.User) *domain.User
	AddBook(userID string, book *domain.Book) (*domain.Book, error)
}

type IUserRepository interface {
	UserReader
	UserWriter
}
