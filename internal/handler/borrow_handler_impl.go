package handler

import (
	"fmt"

	"github.com/nadiannis/libry/internal/usecase"
	"github.com/nadiannis/libry/internal/utils"
)

type BorrowHandler struct {
	usecase usecase.IBorrowUsecase
}

func NewBorrowHandler(usecase usecase.IBorrowUsecase) IBorrowHandler {
	return &BorrowHandler{
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
