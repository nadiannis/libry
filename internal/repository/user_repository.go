package repository

import "github.com/nadiannis/libry/internal/domain"

type UserReader interface {
	GetAllUsers() []*domain.User
	GetUserByUsername(username string) (*domain.User, error)
}

type UserWriter interface {
	AddUser(user *domain.User) *domain.User
}

type IUserRepository interface {
	UserReader
	UserWriter
}
