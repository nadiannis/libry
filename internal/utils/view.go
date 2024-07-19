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
	fmt.Printf("%-40s %-15s %-15s\n", "id", "title", "author")
	fmt.Println(strings.Repeat("-", 80))

	for _, book := range data {
		fmt.Printf("%-40s %-15s %-25s\n", book.ID, book.Title, book.Author)
	}
}

func UserTable(data []*domain.User) {
	fmt.Printf("%-40s %-15s %-25s\n", "id", "username", "books")
	fmt.Println(strings.Repeat("-", 80))

	for _, user := range data {
		fmt.Printf("%-40s %-15s %-15v\n", user.ID, user.Username, user.Books)
	}
}
