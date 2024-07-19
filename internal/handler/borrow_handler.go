package handler

type BorrowReader interface {
	GetAllBorrowedBooks(parts []string)
}

type BorrowWriter interface {
	BorrowBook(parts []string)
}

type IBorrowHandler interface {
	BorrowReader
	BorrowWriter
}
