package usecase

import (
	"github.com/nadiannis/libry/internal/domain"
	"github.com/nadiannis/libry/internal/domain/input"
	"github.com/nadiannis/libry/internal/repository"

	"github.com/google/uuid"
)

type BookUsecase struct {
	repository repository.IBookRepository
}

func NewBookUsecase(repository repository.IBookRepository) IBookUsecase {
	return &BookUsecase{
		repository: repository,
	}
}

func (u *BookUsecase) GetAllBooks() []*domain.Book {
	return u.repository.GetAllBooks()
}

func (u *BookUsecase) AddBook(input *input.BookInput) *domain.Book {
	book := &domain.Book{
		ID:     uuid.NewString(),
		Title:  input.Title,
		Author: input.Author,
	}
	return u.repository.AddBook(book)
}

func (u *BookUsecase) GetBookByID(bookID string) (*domain.Book, error) {
	book, err := u.repository.GetBookByID(bookID)
	if err != nil {
		return nil, err
	}

	return book, nil
}
