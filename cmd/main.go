package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

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
			fmt.Println("List all borrowed books")
		case `\b`:
			fmt.Println("Borrow a book")
		case `\r`:
			fmt.Println("Return a book")
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
	\q   => Quit`)
}
