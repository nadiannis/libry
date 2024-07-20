package usecase

import (
	"github.com/google/uuid"
	"github.com/nadiannis/libry/internal/domain"
	"github.com/nadiannis/libry/internal/domain/input"
	"github.com/nadiannis/libry/internal/repository"
	"github.com/nadiannis/libry/internal/utils"
)

type UserUsecase struct {
	repository repository.IUserRepository
}

func NewUserUsecase(repository repository.IUserRepository) IUserUsecase {
	return &UserUsecase{
		repository: repository,
	}
}

func (u *UserUsecase) GetAllUsers() []*domain.User {
	return u.repository.GetAllUsers()
}

func (u *UserUsecase) AddUser(input *input.UserInput) (*domain.User, error) {
	foundUser, _ := u.repository.GetUserByUsername(input.Username)
	if foundUser != nil {
		return nil, utils.ErrUserExists
	}

	user := &domain.User{
		ID:       uuid.NewString(),
		Username: input.Username,
		Books:    make([]*domain.Book, 0),
	}
	return u.repository.AddUser(user), nil
}

func (u *UserUsecase) GetUserByUsername(username string) (*domain.User, error) {
	user, err := u.repository.GetUserByUsername(username)
	if err != nil {
		return nil, err
	}

	return user, nil
}
