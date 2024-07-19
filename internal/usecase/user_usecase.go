package usecase

import (
	"github.com/nadiannis/libry/internal/domain"
	"github.com/nadiannis/libry/internal/dto"
)

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
