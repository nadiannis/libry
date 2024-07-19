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
