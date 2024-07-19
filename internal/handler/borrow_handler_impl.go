package handler

import (
	"bufio"
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

	fmt.Printf("(book with ID '%s' is borrowed by '%s')\n", borrowedBook.ID, username)
}
