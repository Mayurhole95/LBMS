package user

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/Mayurhole95/LBMS/api"
	"github.com/golang-jwt/jwt"
	"github.com/gorilla/mux"
)

type JWTClaim struct {
	UserID string `json:userid`
	Email  string `json:email`
	Role   string `json:role`
	jwt.StandardClaims
}
type TokenData struct {
	UserID string
	Email  string
	Role   string
}

var jwtKey = []byte("yu78jhe5$r")

func Login() http.HandlerFunc {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		var j JWTClaim
		err := json.NewDecoder(req.Body).Decode(&j)
		if err != nil {
			api.Error(rw, http.StatusBadRequest, api.Response{Message: err.Error()})
			return
		}
		jwtString, err := GenerateJWT(j.UserID, j.Email, j.Role)
		if err != nil {
			api.Error(rw, http.StatusBadRequest, api.Response{Message: err.Error()})
			return
		}

		api.Success(rw, http.StatusCreated, jwtString)
		fmt.Println(jwtString)

	})
}

func GenerateJWT(UserID string, Email string, Role string) (tokenString string, err error) {

	expirationTime := time.Now().Add(1 * time.Hour)
	claims := &JWTClaim{
		Email:  Email,
		UserID: UserID,
		Role:   Role,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err = token.SignedString(jwtKey)
	return
}

func Authorize(handler http.HandlerFunc, role string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		token := r.Header.Get("Authorization")

		isValid, tokenData, err := ValidateToken(token)
		if err != nil {
			fmt.Println("error")
		}

		if !isValid {
			//Send error response to api
		}

		tokenRole := tokenData.Role

		fmt.Println("Role : ", tokenRole)

		handler.ServeHTTP(w, r)
		return
	}
}

func ValidateToken(signedToken string) (isValid bool, tokenData TokenData, err error) {
	token, err := jwt.ParseWithClaims(
		signedToken,
		&JWTClaim{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(jwtKey), nil
		},
	)
	if err != nil {
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
		UserID: claims.UserID,
		Email:  claims.Email,
		Role:   claims.Role,
	}
	return
}

func Create(service Service) http.HandlerFunc {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		var c createRequest
		err := json.NewDecoder(req.Body).Decode(&c)
		if err != nil {
			api.Error(rw, http.StatusBadRequest, api.Response{Message: err.Error()})
			return
		}

		err = service.create(req.Context(), c)
		if isBadRequest(err) {
			api.Error(rw, http.StatusBadRequest, api.Response{Message: err.Error()})
			return
		}

		if err != nil {
			api.Error(rw, http.StatusInternalServerError, api.Response{Message: err.Error()})
			return
		}

		api.Success(rw, http.StatusCreated, api.Response{Message: "Created Successfully"})
	})
}

func List(service Service) http.HandlerFunc {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		resp, err := service.list(req.Context())
		if err == errNoUsers {
			api.Error(rw, http.StatusNotFound, api.Response{Message: err.Error()})
			return
		}
		if err != nil {
			api.Error(rw, http.StatusInternalServerError, api.Response{Message: err.Error()})
			return
		}

		api.Success(rw, http.StatusOK, resp)
	})
}

func FindByID(service Service) http.HandlerFunc {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		vars := mux.Vars(req)

		resp, err := service.findByID(req.Context(), vars["user_id"])

		if err == errNoUserId {
			api.Error(rw, http.StatusNotFound, api.Response{Message: err.Error()})
			return
		}
		if err != nil {
			api.Error(rw, http.StatusInternalServerError, api.Response{Message: err.Error()})
			return
		}

		api.Success(rw, http.StatusOK, resp)
	})
}

func DeleteByID(service Service) http.HandlerFunc {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		vars := mux.Vars(req)

		err := service.deleteByID(req.Context(), vars["user_id"])
		if err == errNoUserId {
			api.Error(rw, http.StatusNotFound, api.Response{Message: err.Error()})
		}
		if err != nil {
			api.Error(rw, http.StatusInternalServerError, api.Response{Message: err.Error()})
			return
		}

		api.Success(rw, http.StatusOK, api.Response{Message: "Deleted Successfully"})
	})
}

func Update(service Service) http.HandlerFunc {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		var c updateRequest
		err := json.NewDecoder(req.Body).Decode(&c)
		if err != nil {
			api.Error(rw, http.StatusBadRequest, api.Response{Message: err.Error()})
			return
		}

		err = service.update(req.Context(), c)
		if isBadRequest(err) {
			api.Error(rw, http.StatusBadRequest, api.Response{Message: err.Error()})
			return
		}

		if err != nil {
			api.Error(rw, http.StatusInternalServerError, api.Response{Message: err.Error()})
			return
		}

		api.Success(rw, http.StatusOK, api.Response{Message: "Updated Successfully"})
	})
}

func isBadRequest(err error) bool {
	return err == errEmptyName || err == errEmptyID
}
