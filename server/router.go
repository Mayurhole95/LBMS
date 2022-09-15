package server

import (
	"net/http"

	"github.com/Mayurhole95/LBMS/api"
	book "github.com/Mayurhole95/LBMS/book"
	"github.com/Mayurhole95/LBMS/transaction"
	user "github.com/Mayurhole95/LBMS/user"
	"github.com/gorilla/mux"
)

const (
	versionHeader = "Accept"
)

func initRouter(dep dependencies) (router *mux.Router) {

	router = mux.NewRouter()
	router.HandleFunc("/ping", pingHandler).Methods(http.MethodGet)

	//User

	router.HandleFunc("/users", middleware(user.Create(dep.UserService), RoleAdmin)).Methods(http.MethodPost)
	router.HandleFunc("/users", middleware(user.List(dep.UserService), RoleAdmin)).Methods(http.MethodGet)
	router.HandleFunc("/users/{user_id}", middleware(user.FindByID(dep.UserService), RoleUser)).Methods(http.MethodGet)
	router.HandleFunc("/users/{user_id}", middleware(user.DeleteByID(dep.UserService), RoleAdmin)).Methods(http.MethodDelete)
	router.HandleFunc("/users", middleware(user.Update(dep.UserService), RoleUser)).Methods(http.MethodPut)

	//Book

	router.HandleFunc("/books", middleware(book.Create(dep.BookService), RoleAdmin)).Methods(http.MethodPost)
	router.HandleFunc("/books", middleware(book.List(dep.BookService), RoleUser)).Methods(http.MethodGet)
	router.HandleFunc("/books/{book_id}", middleware(book.FindByID(dep.BookService), RoleUser)).Methods(http.MethodGet)
	router.HandleFunc("/books/{book_id}", middleware(book.DeleteByID(dep.BookService), Admin)).Methods(http.MethodDelete)
	router.HandleFunc("/books", middleware(book.Update(dep.BookService), RoleAdmin)).Methods(http.MethodPut)

	//Transaction

	router.HandleFunc("/userbook/issue", middleware(transaction.CreateTransaction(dep.TransactionService), RoleAdmin).Methods(http.MethodPost)
	router.HandleFunc("/userbook", middleware(transaction.ListTransaction(dep.TransactionService), RoleAdmin)).Methods(http.MethodGet)
	router.HandleFunc("/userbook/return",middleware(transaction.UpdateTransaction(dep.TransactionService),RoleAdmin)).Methods(http.MethodPut)

	return
}

func pingHandler(rw http.ResponseWriter, req *http.Request) {
	api.Success(rw, http.StatusOK, api.Response{Message: "pong"})
}
