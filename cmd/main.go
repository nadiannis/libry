package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/nadiannis/libry/internal/handler"
	"github.com/nadiannis/libry/internal/repository"
	"github.com/nadiannis/libry/internal/usecase"
	"github.com/nadiannis/libry/internal/utils"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	bookRepo := repository.NewBookRepository()
	bookUsecase := usecase.NewBookUsecase(bookRepo)
	bookHandler := handler.NewBookHandler(bookUsecase)

	userRepo := repository.NewUserRepository()
	userUsecase := usecase.NewUserUsecase(userRepo)
	userHandler := handler.NewUserHandler(userUsecase)

	borrowRepo := repository.NewBorrowRepository()
	borrowUsecase := usecase.NewBorrowUsecase(borrowRepo, userRepo, bookRepo)
	borrowHandler := handler.NewBorrowHandler(scanner, borrowUsecase)

	prepopulateBooks(bookHandler)

	fmt.Println("\n======================= LIBRY =======================")
	displayCommands()

	for {
		input := utils.GetInput(scanner, "\n> ")
		if input == "" {
			fmt.Println("input is required")
			continue
		}

		parts := strings.Fields(input)
		action := parts[0]

		switch action {
		case `\lu`:
			userHandler.GetAllUsers(parts)
		case `\lb`:
			bookHandler.GetAllBooks(parts)
		case `\lbb`:
			borrowHandler.GetAllBorrowedBooks(parts)
		case `\b`:
			borrowHandler.BorrowBook(parts)
		case `\r`:
			borrowHandler.ReturnBook(parts)
		case `\ud`: // to demonstrate if the user is late returning the book
			if len(parts) != 4 {
				fmt.Println(`input should be \ud [borrow-id] [start-date (yyyy-mm-dd)] [end-date (yyyy-mm-dd)]`)
				continue
			}
			startDate, _ := time.Parse("2006-01-02", parts[2])
			endDate, _ := time.Parse("2006-01-02", parts[3])
			borrowRepo.UpdateDates(parts[1], startDate, endDate)
			fmt.Println("borrow dates updated")
		case `\c`:
			if len(parts) != 1 {
				fmt.Println(`input should be \c`)
				continue
			}
			displayCommands()
		case `\q`:
			fmt.Println("bye!")
			return
		default:
			fmt.Println(`action should be '\lu', '\lb', '\lbb', '\b', '\r', or '\q'`)
		}
	}
}

func displayCommands() {
	fmt.Println(`
	Commands:
	\lu  => List all users
	\lb  => List all books
	\lbb => List all borrowed books
	\b   => Borrow a book
	\r   => Return a book
	\c   => Show all commands
	\q   => Quit`)
}
