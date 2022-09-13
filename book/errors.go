package user

import "errors"

var (
	errEmptyID   = errors.New("Book ID must be present")
	errEmptyName = errors.New("Book name must be present")
	errNoBooks   = errors.New("No books present")
	errNoBookId  = errors.New("Book is not present")
)