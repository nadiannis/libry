package utils

import (
	"bufio"
	"fmt"
	"strings"

	"github.com/nadiannis/libry/internal/domain"
)

func GetInput(scanner *bufio.Scanner, prompt string) string {
	fmt.Print(prompt)
	scanner.Scan()
	return strings.TrimSpace(scanner.Text())
}

func BookTable(data []*domain.Book) {
	fmt.Printf("%-40s %-20s %-25s\n", "id", "title", "author")
	fmt.Println(strings.Repeat("-", 85))

	for _, book := range data {
		fmt.Printf("%-40s %-20s %-25s\n", book.ID, book.Title, book.Author)
	}
}

func UserTable(data []*domain.User) {
	for _, user := range data {
		fmt.Printf("%-25s: %-80s\n", "id", user.ID)
		fmt.Printf("%-25s: %-80s\n", "username", fmt.Sprintf("%s (ID: %s)", user.Username, user.ID))

		if len(user.Books) == 0 {
			fmt.Printf("%-25s: %-80s\n", "books currently borrowed", "user does not borrow any books")
		} else {
			for i, book := range user.Books {
				if i == 0 {
					fmt.Printf("%-25s: %-80s\n", "books currently borrowed", fmt.Sprintf("- %s by %s (ID: %s)", book.Title, book.Author, book.ID))
				} else {
					fmt.Printf("%-25s: %-80s\n", "", fmt.Sprintf("- %s by %s (ID: %s)", book.Title, book.Author, book.ID))
				}
			}
		}

		fmt.Println(strings.Repeat("-", 105))
	}
}

func BorrowedBookTable(data []*domain.Borrow) {
	for _, borrowedBook := range data {
		fmt.Printf("%-11s: %-80s\n", "id", borrowedBook.ID)
		fmt.Printf("%-11s: %-80s\n", "book id", borrowedBook.BookID)
		fmt.Printf("%-11s: %-80s\n", "user id", borrowedBook.UserID)
		fmt.Printf("%-11s: %-80s\n", "start date", FormatDate(borrowedBook.StartDate))
		fmt.Printf("%-11s: %-80s\n", "end date", FormatDate(borrowedBook.EndDate))
		fmt.Printf("%-11s: %-80s\n", "status", borrowedBook.Status)
		fmt.Println(strings.Repeat("-", 91))
	}
}
