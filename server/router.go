package server

import (
	"net/http"

	"github.com/Mayurhole95/LBMS/api"
	book "github.com/Mayurhole95/LBMS/book"
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

	router.HandleFunc("/users", user.Create(dep.UserService)).Methods(http.MethodPost)
	router.HandleFunc("/users", user.List(dep.UserService)).Methods(http.MethodGet)
	router.HandleFunc("/users/{user_id}", user.FindByID(dep.UserService)).Methods(http.MethodGet)
	router.HandleFunc("/users/{user_id}", user.DeleteByID(dep.UserService)).Methods(http.MethodDelete)
	router.HandleFunc("/users", user.Update(dep.UserService)).Methods(http.MethodPut)

	//Book

	router.HandleFunc("/books", book.Create(dep.BookService)).Methods(http.MethodPost)
	router.HandleFunc("/books", book.List(dep.BookService)).Methods(http.MethodGet)
	router.HandleFunc("/books/{book_id}", book.FindByID(dep.BookService)).Methods(http.MethodGet)
	router.HandleFunc("/books/{book_id}", book.DeleteByID(dep.BookService)).Methods(http.MethodDelete)
	router.HandleFunc("/books", book.Update(dep.BookService)).Methods(http.MethodPut)

	return
}

func pingHandler(rw http.ResponseWriter, req *http.Request) {
	api.Success(rw, http.StatusOK, api.Response{Message: "pong"})
}
