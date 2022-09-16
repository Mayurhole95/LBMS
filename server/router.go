package server

import (
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/Mayurhole95/LBMS/api"
	book "github.com/Mayurhole95/LBMS/book"
	transaction "github.com/Mayurhole95/LBMS/transaction"
	user "github.com/Mayurhole95/LBMS/user"
	"github.com/golang-jwt/jwt"
	"github.com/gorilla/mux"
)

type JWTClaim struct {
	Id    string `json:"id"`
	Email string `json:"email"`
	Role  string `json:"role"`
	jwt.StandardClaims
}
type TokenData struct {
	Id    string
	Email string
	Role  string
}

const (
	versionHeader = "Accept"
)

const (
	SUPERADMIN = iota
	ADMIN
	USER
)

var RoleMap = map[string]int{"superadmin": SUPERADMIN, "admin": ADMIN, "user": USER}

var jwtKey = []byte("jsd549$^&")

func Authorize(handler http.HandlerFunc, role int) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		//TODO:
		//1. get token from reuqest header
		//2. decode token
		//3. check if user exist from token.ID
		//4. if role is allowed
		//5. call handler

		token := r.Header.Get("Authorization")

		isValid, tokenData, err := ValidateToken(token)
		fmt.Println(isValid)
		if err != nil {
			fmt.Println("error")
		}

		fmt.Println("Token Data : ", tokenData)

		if !isValid {
			//Send error response to api
			api.Error(w, http.StatusBadRequest, api.Response{Message: "Token is not valid"})
			return
		}

		tokenRole := tokenData.Role
		if RoleMap[tokenRole] > role {
			api.Error(w, http.StatusBadRequest, api.Response{Message: "You don't have the access"})
			return
		}

		handler.ServeHTTP(w, r)
		return
	}
}

func ValidateToken(tokenString string) (isValid bool, tokenData TokenData, err error) {
	token, err := jwt.ParseWithClaims(
		tokenString,
		&JWTClaim{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(jwtKey), nil
		},
	)
	if err != nil {
		fmt.Println(err)
		return
	}
	claims, ok := token.Claims.(*JWTClaim)
	if !ok {
		err = errors.New("couldn't parse claims")
		return
	}
	if claims.ExpiresAt < time.Now().Local().Unix() {
		err = errors.New("token expired")
		return
	}

	isValid = true

	tokenData = TokenData{
		Id:    claims.Id,
		Email: claims.Email,
		Role:  claims.Role,
	}
	return
}

func initRouter(dep dependencies) (router *mux.Router) {

	router = mux.NewRouter()
	router.HandleFunc("/ping", pingHandler).Methods(http.MethodGet)

	//User
	router.HandleFunc("/login", user.Login(dep.UserService)).Methods(http.MethodPost)
	router.HandleFunc("/users", Authorize(user.Create(dep.UserService), ADMIN)).Methods(http.MethodPost)
	router.HandleFunc("/users", Authorize(user.List(dep.UserService), ADMIN)).Methods(http.MethodGet)
	router.HandleFunc("/users/{user_id}", Authorize(user.FindByID(dep.UserService), USER)).Methods(http.MethodGet)
	router.HandleFunc("/users/{user_id}", Authorize(user.DeleteByID(dep.UserService), ADMIN)).Methods(http.MethodDelete)
	router.HandleFunc("/users", Authorize(user.Update(dep.UserService), USER)).Methods(http.MethodPut)

	//Book

	router.HandleFunc("/books", Authorize(book.Create(dep.BookService), ADMIN)).Methods(http.MethodPost)
	router.HandleFunc("/books", Authorize(book.List(dep.BookService), USER)).Methods(http.MethodGet)
	router.HandleFunc("/books/{book_id}", Authorize(book.FindByID(dep.BookService), USER)).Methods(http.MethodGet)
	router.HandleFunc("/books/{book_id}", Authorize(book.DeleteByID(dep.BookService), ADMIN)).Methods(http.MethodDelete)
	router.HandleFunc("/books", Authorize(book.Update(dep.BookService), ADMIN)).Methods(http.MethodPut)

	//Transaction

	router.HandleFunc("/userbook/issue", Authorize(transaction.CreateTransaction(dep.TransactionService), ADMIN)).Methods(http.MethodPost)
	router.HandleFunc("/userbook", Authorize(transaction.ListTransaction(dep.TransactionService), ADMIN)).Methods(http.MethodGet)
	router.HandleFunc("/userbook/return", Authorize(transaction.UpdateTransaction(dep.TransactionService), ADMIN)).Methods(http.MethodPut)

	return
}

func pingHandler(rw http.ResponseWriter, req *http.Request) {
	api.Success(rw, http.StatusOK, api.Response{Message: "pong"})
}
