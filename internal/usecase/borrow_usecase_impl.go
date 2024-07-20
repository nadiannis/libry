package usecase

import (
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/nadiannis/libry/internal/domain"
	"github.com/nadiannis/libry/internal/dto"
	"github.com/nadiannis/libry/internal/repository"
)

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

func (u *BorrowUsecase) GetAllBorrowedBooks() []*domain.Borrow {
	return u.borrowRepository.GetAllBorrowedBooks()
}

func (u *BorrowUsecase) BorrowBook(input *dto.BorrowInput) (*domain.Borrow, error) {
	book, err := u.bookRepository.GetBookByID(input.BookID)
	if err != nil {
		return nil, err
	}
	
	user, _ := u.userRepository.GetUserByUsername(input.Username)
	if user == nil {
		newUser := &domain.User{
			ID:       uuid.NewString(),
			Username: input.Username,
			Books:    make([]*domain.Book, 0),
		}
		user = u.userRepository.AddUser(newUser)
		fmt.Printf("(user '%s' is created)\n", user.Username)
	}

	borrow := &domain.Borrow{
		ID:        uuid.NewString(),
		UserID:    user.ID,
		BookID:    book.ID,
		StartDate: time.Now(),
		EndDate:   time.Now().AddDate(0, 0, 7),
	}
	borrowedBook, err := u.borrowRepository.AddBorrowedBook(borrow)
	if err != nil {
		return nil, err
	}

	u.userRepository.AddBook(user.ID, book)

	return borrowedBook, nil
}
