package user_test

import (
	"bytes"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Mayurhole95/LBMS/user"
	usermock "github.com/Mayurhole95/LBMS/user/mocks"
	"github.com/stretchr/testify/mock"
)

func checkResponseCode(t *testing.T, expected, actual int) {
	fmt.Println("Expected Code :", expected, "Actual Code : ", actual)
	if expected != actual {
		t.Errorf("Expected response code %d. Got %d\n", expected, actual)
	}

}

func makeHTTPCall(handler http.HandlerFunc, method, path, body string) (rr *httptest.ResponseRecorder) {
	request := []byte(body)
	req, _ := http.NewRequest(method, path, bytes.NewBuffer(request))
	rr = httptest.NewRecorder()
	handler.ServeHTTP(rr, req)
	return
}

// Create:
func TestSuccessfullCreate(t *testing.T) {
	cs := &usermock.Service{}
	cs.On("Create", mock.Anything, mock.Anything).Return(nil, errors.New("Success "))

	rr := makeHTTPCall(user.Create(cs), http.MethodPost, "/users", `{
        "id":"67",
        "first_name": "Rushikesh",
        "last_name": "Markad",
        "gender": "Male",
        "age": 22,
        "address": "Pune",
        "email": "markaaa@gmail.com",
        "password": "ertikol@123",
        "mob_no": "6985749623",
        "role":"user"
    }`)
	fmt.Println("Account Success")
	checkResponseCode(t, http.StatusCreated, rr.Code)
	cs.AssertExpectations(t)
}

func TestCreateWhenInvalidRequestBody(t *testing.T) {
	cs := &usermock.Service{}
	// cs.On("Create", mock.Anything, mock.Anything).Return(nil, errors.New("HEllo"))
	rr := makeHTTPCall(user.Create(cs), http.MethodPost, "/users", `{
        "id":"67",
        "first_name": "Rushikesh",
        "last_name": "Markad",
        "gender": "Male",
        "age": 22,
        "address": "Pune",
        "email": "markaaa@gmail.com",
        "password": "ertikol@123",
        "mob_no": 6985749623,
        "role":"user"
    }`)

	checkResponseCode(t, http.StatusBadRequest, rr.Code)
	cs.AssertExpectations(t)
}

func TestCreateWhenEmptyName(t *testing.T) {
	cs := &usermock.Service{}
	//cs.On("Create", mock.Anything, mock.Anything).Return(nil, errors.New("Empty Name"))

	rr := makeHTTPCall(user.Create(cs), http.MethodPost, "/users", `{
        "id":"67",
        "first_name": "",
        "last_name": "hole",
        "gender": "Male",
        "age": 22,
        "address": "Pune",
        "email": "mh@gmail.com",
        "password": "ertikol@123",
        "mob_no": "6985749623",
        "role":"user"
    }`)

	checkResponseCode(t, http.StatusBadRequest, rr.Code)
	cs.AssertExpectations(t)
}

func TestSuccessfullList(t *testing.T) {
	var lr user.ListResponse
	cs := &usermock.Service{}
	cs.On("List", mock.Anything).Return(lr, nil)

	rr := makeHTTPCall(user.List(cs), http.MethodGet, "/users", "")

	checkResponseCode(t, http.StatusOK, rr.Code)
	cs.AssertExpectations(t)
}
func TestListInternalError(t *testing.T) {
	cs := &usermock.Service{}
	var lr user.ListResponse
	cs.On("List", mock.Anything).Return(lr, errors.New("Internal Error"))

	rr := makeHTTPCall(user.List(cs), http.MethodGet, "/users", "")

	checkResponseCode(t, http.StatusInternalServerError, rr.Code)
	cs.AssertExpectations(t)
}

func TestSuccessfullFindByID(t *testing.T) {
	cs := &usermock.Service{}
	var lr user.FindByIDResponse
	cs.On("FindByID", mock.Anything, mock.Anything).Return(lr, nil)

	rr := makeHTTPCall(user.FindByID(cs), http.MethodGet, "/users/f9c0597f-3069-463b-bfed-28e306df181d", "")

	checkResponseCode(t, http.StatusOK, rr.Code)
	cs.AssertExpectations(t)
}
func TestFindByIDWhenIDNotExist(t *testing.T) {
	cs := &usermock.Service{}
	var lr user.FindByIDResponse
	cs.On("FindByID", mock.Anything, mock.Anything).Return(lr, errors.New("err"))

	rr := makeHTTPCall(user.FindByID(cs), http.MethodGet, "/users/035cb200-2619-4df8-8d2d-9661fc204e55", "")

	checkResponseCode(t, http.StatusNotFound, rr.Code)
	cs.AssertExpectations(t)
}

func TestFindByIDWhenInternalError(t *testing.T) {
	cs := &usermock.Service{}
	var lr user.FindByIDResponse
	cs.On("FindByID", mock.Anything, mock.Anything).Return(lr, errors.New("err"))

	rr := makeHTTPCall(user.FindByID(cs), http.MethodGet, "/users/f9c0597f-3069-463b-bfed-28e306df182q", "")

	checkResponseCode(t, http.StatusInternalServerError, rr.Code)
	cs.AssertExpectations(t)
}

func TestSuccessfullDeleteByID(t *testing.T) {
	cs := &usermock.Service{}
	cs.On("DeleteByID", mock.Anything, mock.Anything).Return(nil)

	rr := makeHTTPCall(user.DeleteByID(cs), http.MethodDelete, "/users/", "f9c0597f-3069-463b-bfed-28e306df181d")

	checkResponseCode(t, http.StatusOK, rr.Code)
	cs.AssertExpectations(t)
}

func TestDeleteByIDWhenIDNotExist(t *testing.T) {
	cs := &usermock.Service{}
	cs.On("DeleteByID", mock.Anything, mock.Anything).Return("error")

	rr := makeHTTPCall(user.DeleteByID(cs), http.MethodDelete, "/users/f9c0597f-3069-463b-bfed-28e306df182q", "")

	checkResponseCode(t, http.StatusNotFound, rr.Code)
	cs.AssertExpectations(t)
}

func TestDeleteByIDWhenInternalError(t *testing.T) {
	cs := &usermock.Service{}
	cs.On("DeleteByID", mock.Anything, mock.Anything).Return(errors.New("Internal Error"))

	rr := makeHTTPCall(user.DeleteByID(cs), http.MethodDelete, "/users/83b34ad3-5803-47ce-b10e-1f9a940eb78", "")

	checkResponseCode(t, http.StatusInternalServerError, rr.Code)
	cs.AssertExpectations(t)
}

func TestSuccessfullUpdate(t *testing.T) {
	cs := &usermock.Service{}
	cs.On("Update", mock.Anything, mock.Anything).Return(nil)

	rr := makeHTTPCall(user.Update(cs), http.MethodPut, "/users", `{"id":"83b34ad3-5803-47ce-b10e-1f9a940eb78", "name":"omkar"}`)

	checkResponseCode(t, http.StatusOK, rr.Code)
	cs.AssertExpectations(t)
}

func TestUpdateWhenInvalidRequestBody(t *testing.T) {
	cs := &usermock.Service{}

	rr := makeHTTPCall(user.Update(cs), http.MethodPut, "/users", `{"id":"83b34ad3-5803-47ce-b10e-1f9a940eb78", "name":"omkar",}`)

	checkResponseCode(t, http.StatusBadRequest, rr.Code)
	cs.AssertExpectations(t)
}
