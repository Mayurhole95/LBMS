package db

import "errors"

var (
	ErrUserNotExist = errors.New("User does not exist in db")
	ErrBookNotExist = errors.New("Book does not exist in db")
)
