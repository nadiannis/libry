<div align="center">
  <br>
  <h1>Libry</h1>
  <p>📖 A simple library CLI app 📖</p>
  <br>
</div>

## Table of Contents

- [Description](#description)
- [Features](#features)
- [Entities](#entities)
- [Usage](#usage)
- [Which Parts Apply SOLID](#which-parts-apply-solid)

## Description

[`^ back to top ^`](#table-of-contents)

**Libry** is a simple CLI app that allows people to borrow books from a library. User of the app can be created when borrowing a book. The app is made with Go. It is created as a submission for the first-week exam in the Go phase of the Backend Development Training.

## Features

[`^ back to top ^`](#table-of-contents)

- View list of users & their currently borrowed books.
- View list of books.
- View list of borrowed books.
- Borrow a book.
- Return a book.
- Show all commands.
- Quit the app.

## Entities

[`^ back to top ^`](#table-of-contents)

There are 4 entities: **Book**, **User**, **Borrow**, & **BorrowStatus**.

**Book**

- id: `string`
- title: `string`
- author: `string`

**User**

- id: `string`
- username: `string`
- books: `[]*Book`

**Borrow**

- id: `string`
- book_id: `string`
- user_id: `string`
- start_date: `timestamp`
- end_date: `timestamp`
- BorrowStatus

**BorrowStatus**

- status: `string`

## Usage

[`^ back to top ^`](#table-of-contents)

- Start the app.

  ```bash
  go run ./cmd
  ```

- Enter a command.

  Choose one of these commands:

  **View list of users & their currently borrowed books**

  ```bash
  \lu
  ```

  **View list of books**

  ```bash
  \lb
  ```

  **View list of borrowed books**

  ```bash
  \lbb
  ```

  **Borrow a book**

  ```bash
  \b
  ```

  **Return a book**

  ```bash
  \r
  ```

  **Show all commands**

  ```bash
  \c
  ```

  **Quit the app**

  ```bash
  \q
  ```

## Which Parts Apply SOLID

[`^ back to top ^`](#table-of-contents)

This project is also intended to learn and apply SOLID principles. The following are the parts that apply those principles:

### Single Responsibility Principle (SRP)

[`^ back to top ^`](#table-of-contents)

```go
type Book struct {
	ID     string
	Title  string
	Author string
}

type User struct {
	ID       string
	Username string
	Books    []*Book
}

type BorrowStatus struct {
	Status string
}

type Borrow struct {
	ID        string
	BookID    string
	UserID    string
	StartDate time.Time
	EndDate   time.Time
	BorrowStatus
}
```

```go
func FormatDate(datetime time.Time) string {
	format := "02 Jan 2006"
	return datetime.Format(format)
}

func TimeIsBetween(t, min, max time.Time) bool {
	if min.After(max) {
		min, max = max, min
	}
	return (t.Equal(min) || t.After(min)) && (t.Equal(max) || t.Before(max))
}

func DateIsAfter(t1, t2 time.Time) bool {
	date1 := t1.Truncate(24 * time.Hour)
	date2 := t2.Truncate(24 * time.Hour)
	return date1.After(date2)
}

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
```

```go
package handler

import "github.com/nadiannis/libry/internal/domain/input"

type BookReader interface {
	GetAllBooks(parts []string)
}

type BookWriter interface {
	AddBook(input *input.BookInput)
}

type IBookHandler interface {
	BookReader
	BookWriter
}

type UserReader interface {
	GetAllUsers(parts []string)
}

type UserWriter interface {
	AddUser(input *input.UserInput)
}

type IUserHandler interface {
	UserReader
	UserWriter
}

type BorrowReader interface {
	GetAllBorrowedBooks(parts []string)
}

type BorrowWriter interface {
	BorrowBook(parts []string)
	ReturnBook(parts []string)
}

type IBorrowHandler interface {
	BorrowReader
	BorrowWriter
}
```

```go
package usecase

import (
	"github.com/nadiannis/libry/internal/domain"
	"github.com/nadiannis/libry/internal/domain/input"
)

type BookReader interface {
	GetAllBooks() []*domain.Book
	GetBookByID(bookID string) (*domain.Book, error)
}

type BookWriter interface {
	AddBook(book *input.BookInput) *domain.Book
}

type IBookUsecase interface {
	BookReader
	BookWriter
}

type UserReader interface {
	GetAllUsers() []*domain.User
	GetUserByUsername(username string) (*domain.User, error)
}

type UserWriter interface {
	AddUser(user *input.UserInput) (*domain.User, error)
}

type IUserUsecase interface {
	UserReader
	UserWriter
}

type BorrowReader interface {
	GetAllBorrowedBooks() []*domain.Borrow
}

type BorrowWriter interface {
	BorrowBook(input *input.BorrowInput) (*domain.Borrow, error)
	ReturnBook(input *input.BorrowInput) (*domain.Borrow, error)
}

type IBorrowUsecase interface {
	BorrowReader
	BorrowWriter
}
```

```go
package repository

import (
	"time"

	"github.com/nadiannis/libry/internal/domain"
)

type BookReader interface {
	GetAllBooks() []*domain.Book
	GetBookByID(bookID string) (*domain.Book, error)
}

type BookWriter interface {
	AddBook(book *domain.Book) *domain.Book
}

type IBookRepository interface {
	BookReader
	BookWriter
}

type UserReader interface {
	GetAllUsers() []*domain.User
	GetUserByUsername(username string) (*domain.User, error)
}

type UserWriter interface {
	AddUser(user *domain.User) *domain.User
	AddBook(userID string, book *domain.Book) (*domain.Book, error)
	DeleteBookByID(userID, bookID string) error
}

type IUserRepository interface {
	UserReader
	UserWriter
}

type BorrowReader interface {
	GetAllBorrowedBooks() []*domain.Borrow
	GetBorrowedBook(userID, bookID string) (*domain.Borrow, error)
}

type BorrowWriter interface {
	AddBorrowedBook(borrow *domain.Borrow) (*domain.Borrow, error)
	UpdateStatus(borrowID, status string) (*domain.Borrow, error)
	UpdateDates(borrowID string, startDate, endDate time.Time) error
}

type IBorrowRepository interface {
	BorrowReader
	BorrowWriter
}
```

### Open-Closed Principle (OCP)

[`^ back to top ^`](#table-of-contents)

```go
type UserHandler struct {
	usecase usecase.IUserUsecase
}

func NewUserHandler(usecase usecase.IUserUsecase) IUserHandler {
	return &UserHandler{
		usecase: usecase,
	}
}
```

```go
type BookHandler struct {
	usecase usecase.IBookUsecase
}

func NewBookHandler(usecase usecase.IBookUsecase) IBookHandler {
	return &BookHandler{
		usecase: usecase,
	}
}
```

```go
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
```

```go
type UserUsecase struct {
	repository repository.IUserRepository
}

func NewUserUsecase(repository repository.IUserRepository) IUserUsecase {
	return &UserUsecase{
		repository: repository,
	}
}
```

```go
type BookUsecase struct {
	repository repository.IBookRepository
}

func NewBookUsecase(repository repository.IBookRepository) IBookUsecase {
	return &BookUsecase{
		repository: repository,
	}
}
```

```go
type BorrowUsecase struct {
	borrowRepository repository.IBorrowRepository
	userRepository   repository.IUserRepository
	bookRepository   repository.IBookRepository
}

func NewBorrowUsecase(
	borrowRepository repository.IBorrowRepository,
	userRepository repository.IUserRepository,
	bookRepository repository.IBookRepository,
) IBorrowUsecase {
	return &BorrowUsecase{
		borrowRepository: borrowRepository,
		userRepository:   userRepository,
		bookRepository:   bookRepository,
	}
}
```

### Liskov Substitution Principle (LSP)

[`^ back to top ^`](#table-of-contents)

```go
func NewUserHandler(usecase usecase.IUserUsecase) IUserHandler {
	return &UserHandler{
		usecase: usecase,
	}
}
```

```go
func NewBookHandler(usecase usecase.IBookUsecase) IBookHandler {
	return &BookHandler{
		usecase: usecase,
	}
}
```

```go
func NewBorrowHandler(scanner *bufio.Scanner, usecase usecase.IBorrowUsecase) IBorrowHandler {
	return &BorrowHandler{
		scanner: scanner,
		usecase: usecase,
	}
}
```

```go
func NewUserUsecase(repository repository.IUserRepository) IUserUsecase {
	return &UserUsecase{
		repository: repository,
	}
}
```

```go
func NewBookUsecase(repository repository.IBookRepository) IBookUsecase {
	return &BookUsecase{
		repository: repository,
	}
}
```

```go
func NewBorrowUsecase(
	borrowRepository repository.IBorrowRepository,
	userRepository repository.IUserRepository,
	bookRepository repository.IBookRepository,
) IBorrowUsecase {
	return &BorrowUsecase{
		borrowRepository: borrowRepository,
		userRepository:   userRepository,
		bookRepository:   bookRepository,
	}
}
```

```go
func NewUserRepository() IUserRepository {
	return &UserRepository{
		db: make(map[string]*domain.User),
	}
}
```

```go
func NewBookRepository() IBookRepository {
	return &BookRepository{
		db: make(map[string]*domain.Book),
	}
}
```

```go
func NewBorrowRepository() IBorrowRepository {
	return &BorrowRepository{
		db: make(map[string]*domain.Borrow),
	}
}
```

### Interface Segregation Principle (ISP)

[`^ back to top ^`](#table-of-contents)

```go
package handler

import "github.com/nadiannis/libry/internal/domain/input"

type BookReader interface {
	GetAllBooks(parts []string)
}

type BookWriter interface {
	AddBook(input *input.BookInput)
}

type IBookHandler interface {
	BookReader
	BookWriter
}

type UserReader interface {
	GetAllUsers(parts []string)
}

type UserWriter interface {
	AddUser(input *input.UserInput)
}

type IUserHandler interface {
	UserReader
	UserWriter
}

type BorrowReader interface {
	GetAllBorrowedBooks(parts []string)
}

type BorrowWriter interface {
	BorrowBook(parts []string)
	ReturnBook(parts []string)
}

type IBorrowHandler interface {
	BorrowReader
	BorrowWriter
}
```

```go
package usecase

import (
	"github.com/nadiannis/libry/internal/domain"
	"github.com/nadiannis/libry/internal/domain/input"
)

type BookReader interface {
	GetAllBooks() []*domain.Book
	GetBookByID(bookID string) (*domain.Book, error)
}

type BookWriter interface {
	AddBook(book *input.BookInput) *domain.Book
}

type IBookUsecase interface {
	BookReader
	BookWriter
}

type UserReader interface {
	GetAllUsers() []*domain.User
	GetUserByUsername(username string) (*domain.User, error)
}

type UserWriter interface {
	AddUser(user *input.UserInput) (*domain.User, error)
}

type IUserUsecase interface {
	UserReader
	UserWriter
}

type BorrowReader interface {
	GetAllBorrowedBooks() []*domain.Borrow
}

type BorrowWriter interface {
	BorrowBook(input *input.BorrowInput) (*domain.Borrow, error)
	ReturnBook(input *input.BorrowInput) (*domain.Borrow, error)
}

type IBorrowUsecase interface {
	BorrowReader
	BorrowWriter
}
```

```go
package repository

import (
	"time"

	"github.com/nadiannis/libry/internal/domain"
)

type BookReader interface {
	GetAllBooks() []*domain.Book
	GetBookByID(bookID string) (*domain.Book, error)
}

type BookWriter interface {
	AddBook(book *domain.Book) *domain.Book
}

type IBookRepository interface {
	BookReader
	BookWriter
}

type UserReader interface {
	GetAllUsers() []*domain.User
	GetUserByUsername(username string) (*domain.User, error)
}

type UserWriter interface {
	AddUser(user *domain.User) *domain.User
	AddBook(userID string, book *domain.Book) (*domain.Book, error)
	DeleteBookByID(userID, bookID string) error
}

type IUserRepository interface {
	UserReader
	UserWriter
}

type BorrowReader interface {
	GetAllBorrowedBooks() []*domain.Borrow
	GetBorrowedBook(userID, bookID string) (*domain.Borrow, error)
}

type BorrowWriter interface {
	AddBorrowedBook(borrow *domain.Borrow) (*domain.Borrow, error)
	UpdateStatus(borrowID, status string) (*domain.Borrow, error)
	UpdateDates(borrowID string, startDate, endDate time.Time) error
}

type IBorrowRepository interface {
	BorrowReader
	BorrowWriter
}
```

### Dependency Inversion Principle (DIP)

[`^ back to top ^`](#table-of-contents)

```go
type IUserUsecase interface {
	UserReader
	UserWriter
}

type IUserHandler interface {
	UserReader
	UserWriter
}

type UserHandler struct {
	usecase usecase.IUserUsecase
}

func NewUserHandler(usecase usecase.IUserUsecase) IUserHandler {
	return &UserHandler{
		usecase: usecase,
	}
}
```

```go
type IBookUsecase interface {
	BookReader
	BookWriter
}

type IBookHandler interface {
	BookReader
	BookWriter
}

type BookHandler struct {
	usecase usecase.IBookUsecase
}

func NewBookHandler(usecase usecase.IBookUsecase) IBookHandler {
	return &BookHandler{
		usecase: usecase,
	}
}
```

```go
type IBorrowUsecase interface {
	BorrowReader
	BorrowWriter
}

type IBorrowHandler interface {
	BorrowReader
	BorrowWriter
}

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
```

```go
type IUserRepository interface {
	UserReader
	UserWriter
}

type IUserUsecase interface {
	UserReader
	UserWriter
}

type UserUsecase struct {
	repository repository.IUserRepository
}

func NewUserUsecase(repository repository.IUserRepository) IUserUsecase {
	return &UserUsecase{
		repository: repository,
	}
}
```

```go
type IBookRepository interface {
	BookReader
	BookWriter
}

type IBookUsecase interface {
	BookReader
	BookWriter
}

type BookUsecase struct {
	repository repository.IBookRepository
}

func NewBookUsecase(repository repository.IBookRepository) IBookUsecase {
	return &BookUsecase{
		repository: repository,
	}
}
```

```go
type IUserRepository interface {
	UserReader
	UserWriter
}

type UserRepository struct {
	db map[string]*domain.User
}

func NewUserRepository() IUserRepository {
	return &UserRepository{
		db: make(map[string]*domain.User),
	}
}
```

```go
type IBookRepository interface {
	BookReader
	BookWriter
}

type BookRepository struct {
	db map[string]*domain.Book
}

func NewBookRepository() IBookRepository {
	return &BookRepository{
		db: make(map[string]*domain.Book),
	}
}
```

```go
type IBorrowRepository interface {
	BorrowReader
	BorrowWriter
}

type BorrowRepository struct {
	db map[string]*domain.Borrow
}

func NewBorrowRepository() IBorrowRepository {
	return &BorrowRepository{
		db: make(map[string]*domain.Borrow),
	}
}
```
