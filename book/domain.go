package book

import "github.com/Mayurhole95/LBMS/db"

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
	return

	if cr.Status == "" || cr.Status != "Available" && cr.Status != "Unavailable" {
		return errInvalidStatus
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
