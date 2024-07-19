package main

import (
	"github.com/nadiannis/libry/internal/dto"
	"github.com/nadiannis/libry/internal/handler"
)

var bookInputs = []*dto.BookInput{
	{
		Title:  "Book 1",
		Author: "Author 1",
	},
	{
		Title:  "Book 2",
		Author: "Author 2",
	},
	{
		Title:  "Book 3",
		Author: "Author 3",
	},
	{
		Title:  "Book 4",
		Author: "Author 4",
	},
	{
		Title:  "Book 5",
		Author: "Author 5",
	},
}

func prepopulateBooks(bookHandler handler.IBookHandler) {
	for _, bookInput := range bookInputs {
		bookHandler.AddBook(bookInput)
	}
}
