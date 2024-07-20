package handler

import (
	"bufio"
	"errors"
	"fmt"

	"github.com/nadiannis/libry/internal/dto"
	"github.com/nadiannis/libry/internal/usecase"
	"github.com/nadiannis/libry/internal/utils"
)

type BorrowHandler struct {
	scanner *bufio.Scanner
	usecase usecase.IBorrowUsecase
}

func NewBorrowHandler(scanner *bufio.Scanner, usecase usecase.IBorrowUsecase) IBorrowHandler {
	return &BorrowHandler{
		scanner: scanner,
		usecase: usecase,
	}
}

func (h *BorrowHandler) GetAllBorrowedBooks(parts []string) {
	if len(parts) != 1 {
		fmt.Println(`input should be \lbb`)
		return
	}

	borrowedBooks := h.usecase.GetAllBorrowedBooks()
	total := len(borrowedBooks)

	fmt.Printf("total: %d\n", total)

	if total == 0 {
		fmt.Println("there are no books borrowed")
		return
	}

	utils.BorrowedBookTable(borrowedBooks)
}

func (h *BorrowHandler) BorrowBook(parts []string) {
	if len(parts) != 1 {
		fmt.Println(`input should be \b`)
		return
	}

	var username, bookID string
	var err error

	for {
		username = utils.GetInput(h.scanner, ">> Enter username: ")
		if username == "" {
			fmt.Println("username is required")
			continue
		}
		break
	}

	for {
		bookID = utils.GetInput(h.scanner, ">> Enter book ID: ")
		if bookID == "" {
			fmt.Println("book ID is required")
			continue
		}
		break
	}

	borrowInput := &dto.BorrowInput{
		Username: username,
		BookID:   bookID,
	}
	borrowedBook, err := h.usecase.BorrowBook(borrowInput)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("(book with ID '%s' is borrowed by '%s')\n", borrowedBook.BookID, username)
}

func (h *BorrowHandler) ReturnBook(parts []string) {
	if len(parts) != 1 {
		fmt.Println(`input should be \r`)
		return
	}

	var username, bookID string
	var err error

	for {
		username = utils.GetInput(h.scanner, ">> Enter username: ")
		if username == "" {
			fmt.Println("username is required")
			continue
		}
		break
	}

	for {
		bookID = utils.GetInput(h.scanner, ">> Enter book ID: ")
		if bookID == "" {
			fmt.Println("book ID is required")
			continue
		}
		break
	}

	borrowInput := &dto.BorrowInput{
		Username: username,
		BookID:   bookID,
	}
	returnedBook, err := h.usecase.ReturnBook(borrowInput)
	if err != nil {
		switch {
		case errors.Is(err, utils.ErrOverdueBookReturned):
			fmt.Printf("(book with ID '%s' is returned late by '%s')\n", returnedBook.BookID, username)
		default:
			fmt.Println(err)
		}
		return
	}

	fmt.Printf("(book with ID '%s' is returned by '%s')\n", returnedBook.BookID, username)
}
