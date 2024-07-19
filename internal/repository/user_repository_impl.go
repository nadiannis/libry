package repository

import (
	"github.com/nadiannis/libry/internal/domain"
	"github.com/nadiannis/libry/internal/utils"
)

type UserRepository struct {
	db map[string]*domain.User
}

func NewUserRepository() IUserRepository {
	return &UserRepository{
		db: make(map[string]*domain.User),
	}
}

func (r *UserRepository) GetAllUsers() []*domain.User {
	users := make([]*domain.User, 0)
	for _, user := range r.db {
		users = append(users, user)
	}
	return users
}

func (r *UserRepository) AddUser(user *domain.User) *domain.User {
	r.db[user.ID] = user
	return user
}

func (r *UserRepository) GetUserByUsername(username string) (*domain.User, error) {
	user, exists := r.db[username]
	if !exists {
		return nil, utils.ErrUserNotFound
	}

	return user, nil
}
