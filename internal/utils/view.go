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
	fmt.Printf("%-40s %-15s %-15s\n", "ID", "Title", "Author")
	fmt.Println(strings.Repeat("-", 70))

	for _, book := range data {
		fmt.Printf("%-40s %-15s %-15s\n", book.ID, book.Title, book.Author)
	}
}
