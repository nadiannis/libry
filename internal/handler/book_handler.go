package handler

import (
	"fmt"

	"github.com/nadiannis/libry/internal/domain/input"
	"github.com/nadiannis/libry/internal/usecase"
	"github.com/nadiannis/libry/internal/utils"
)

type BookHandler struct {
	usecase usecase.IBookUsecase
}

func NewBookHandler(usecase usecase.IBookUsecase) IBookHandler {
	return &BookHandler{
		usecase: usecase,
	}
}

func (h *BookHandler) GetAllBooks(parts []string) {
	if len(parts) != 1 {
		fmt.Println(`input should be \lb`)
		return
	}

	books := h.usecase.GetAllBooks()
	total := len(books)

	fmt.Printf("total: %d\n", total)

	if total == 0 {
		fmt.Println("there are no books available")
		return
	}

	utils.BookTable(books)
}

func (h *BookHandler) AddBook(input *input.BookInput) {
	if input.Title == "" {
		fmt.Println("title is required")
		return
	}

	if input.Author == "" {
		fmt.Println("author is required")
		return
	}

	savedBook := h.usecase.AddBook(input)
	fmt.Printf("('%s' by '%s' is saved with ID '%s')\n",
		savedBook.Title, savedBook.Author, savedBook.ID)
}
