package handler

import "github.com/nadiannis/libry/internal/dto"

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
