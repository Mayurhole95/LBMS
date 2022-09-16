package server

import (
	"net/http"

	"github.com/Mayurhole95/LBMS/api"
	book "github.com/Mayurhole95/LBMS/book"
	transaction "github.com/Mayurhole95/LBMS/transaction"
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

	router.HandleFunc("/users", user.Authorize(user.Create(dep.UserService), "RoleAdmin")).Methods(http.MethodPost)
	router.HandleFunc("/login", user.Login()).Methods(http.MethodPost)
	router.HandleFunc("/users", user.Authorize(user.List(dep.UserService), "RoleAdmin")).Methods(http.MethodGet)
	router.HandleFunc("/users/{user_id}", user.Authorize(user.FindByID(dep.UserService), "RoleUser")).Methods(http.MethodGet)
	router.HandleFunc("/users/{user_id}", user.Authorize(user.DeleteByID(dep.UserService), "RoleAdmin")).Methods(http.MethodDelete)
	router.HandleFunc("/users", user.Authorize(user.Update(dep.UserService), "RoleUser")).Methods(http.MethodPut)

	//Book

	router.HandleFunc("/books", user.Authorize(book.Create(dep.BookService), "RoleAdmin")).Methods(http.MethodPost)
	router.HandleFunc("/books", user.Authorize(book.List(dep.BookService), "RoleUser")).Methods(http.MethodGet)
	router.HandleFunc("/books/{book_id}", user.Authorize(book.FindByID(dep.BookService), "RoleUser")).Methods(http.MethodGet)
	router.HandleFunc("/books/{book_id}", user.Authorize(book.DeleteByID(dep.BookService), "Admin")).Methods(http.MethodDelete)
	router.HandleFunc("/books", user.Authorize(book.Update(dep.BookService), "RoleAdmin")).Methods(http.MethodPut)

	//Transaction

	router.HandleFunc("/userbook/issue", user.Authorize(transaction.CreateTransaction(dep.TransactionService), "RoleAdmin")).Methods(http.MethodPost)
	router.HandleFunc("/userbook", user.Authorize(transaction.ListTransaction(dep.TransactionService), "RoleAdmin")).Methods(http.MethodGet)
	router.HandleFunc("/userbook/return", user.Authorize(transaction.UpdateTransaction(dep.TransactionService), "RoleAdmin")).Methods(http.MethodPut)

	return
}

func pingHandler(rw http.ResponseWriter, req *http.Request) {
	api.Success(rw, http.StatusOK, api.Response{Message: "pong"})
}
