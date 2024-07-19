package handler

type BorrowReader interface {
	GetAllBorrowedBooks(parts []string)
}

type IBorrowHandler interface {
	BorrowReader
}
