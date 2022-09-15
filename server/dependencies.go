package server

import (
	"github.com/Mayurhole95/LBMS/app"
	book "github.com/Mayurhole95/LBMS/book"
	"github.com/Mayurhole95/LBMS/db"
	"github.com/Mayurhole95/LBMS/transaction"
	"github.com/Mayurhole95/LBMS/user"
)

type dependencies struct {
	UserService        user.Service
	BookService        book.Service
	TransactionService transaction.Service
}

func initDependencies() (dependencies, error) {
	appDB := app.GetDB()
	logger := app.GetLogger()
	dbStore := db.NewStorer(appDB)

	userService := user.NewService(dbStore, logger)
	bookService := book.NewService(dbStore, logger)
	transactionService := transaction.NewService(dbStore, logger)

	return dependencies{
		UserService:        userService,
		BookService:        bookService,
		TransactionService: transactionService,
	}, nil
}
