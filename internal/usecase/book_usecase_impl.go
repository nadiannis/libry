package usecase

import (
	"github.com/nadiannis/libry/internal/domain"
	"github.com/nadiannis/libry/internal/dto"
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

func (r *BookUsecase) GetAllBooks() []*domain.Book {
	return r.repository.GetAllBooks()
}

func (r *BookUsecase) AddBook(input *dto.BookInput) *domain.Book {
	book := &domain.Book{
		ID:     uuid.NewString(),
		Title:  input.Title,
		Author: input.Author,
	}
	return r.repository.AddBook(book)
}

func (r *BookUsecase) GetBookByID(bookID string) (*domain.Book, error) {
	book, err := r.repository.GetBookByID(bookID)
	if err != nil {
		return nil, err
	}

	return book, nil
}
