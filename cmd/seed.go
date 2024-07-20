package main

import (
	"github.com/nadiannis/libry/internal/domain/input"
	"github.com/nadiannis/libry/internal/handler"
)

var bookInputs = []*input.BookInput{
	{
		Title:  "Grit",
		Author: "Angela Duckworth",
	},
	{
		Title:  "Mindset",
		Author: "Carol Dweck",
	},
	{
		Title:  "Ikigai",
		Author: "Hector Garcia",
	},
	{
		Title:  "Drive",
		Author: "Daniel H. Pink",
	},
	{
		Title:  "Range",
		Author: "David Epstein",
	},
	{
		Title:  "Sapiens",
		Author: "Yuval N. Harari",
	},
	{
		Title:  "Essentialism",
		Author: "Greg McKeown",
	},
	{
		Title:  "Deep Work",
		Author: "Cal Newport",
	},
	{
		Title:  "The Lean Startup",
		Author: "Eric Ries",
	},
	{
		Title:  "Atomic Habits",
		Author: "James Clear",
	},
}

func prepopulateBooks(bookHandler handler.IBookHandler) {
	for _, bookInput := range bookInputs {
		bookHandler.AddBook(bookInput)
	}
}
