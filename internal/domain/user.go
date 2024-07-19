package domain

type User struct {
	ID       string
	Name     string
	Password string
	Books    []Book
}
