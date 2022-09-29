package book

import (
	"unicode"

	"github.com/Mayurhole95/LBMS/db"
)

type UpdateRequest struct {
	ID              string `json:"id"`
	Name            string `json:"name"`
	Author          string `json:"author"`
	Price           int    `json:"price"`
	TotalCopies     int    `json:"total_copies"`
	Status          string `json:"status"`
	AvailableCopies int    `json:"available_copies"`
}

type CreateRequest struct {
	ID              string `json:"id"`
	Name            string `json:"name"`
	Author          string `json:"author"`
	Price           int    `json:"price"`
	TotalCopies     int    `json:"total_copies"`
	Status          string `json:"status"`
	AvailableCopies int    `json:"available_copies"`
}

type findByIDResponse struct {
	Book db.Book `json:"book"`
}

type listResponse struct {
	Books []db.Book `json:"books"`
}

func (cr CreateRequest) Validate() (err error) {
	if cr.Name == "" {
		return errEmptyName
	}
	if cr.Author == "" {
		return errEmptyAuthor
	}

	if cr.TotalCopies == 0 {
		return errZeroCopies
	}
	if !unicode.IsNumber(rune(cr.TotalCopies)) {
		return errInvalidTotalCopies
	}
	if cr.Price < 1 {
		return errInvalidPrice
	}
	// if !unicode.IsNumber(cr.Price) {
	//  return errInvalidPrice
	// }
	if !unicode.IsNumber(rune(cr.Price)) {
		return errInvalidPrice
	}

	if cr.Status != "available" {
		return errInvalidStatus
	}
	if cr.AvailableCopies > cr.TotalCopies {
		return errInvalidAvailableCopies
	}
	if !unicode.IsNumber(rune(cr.AvailableCopies)) {
		return err1InvalidAvailableCopies
	}
	return
}

func (ur UpdateRequest) Validate() (err error) {
	if ur.ID == "" {
		return errEmptyID
	}
	if ur.Name == "" {
		return errEmptyName
	}
	return
}
