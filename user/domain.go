package user

import "github.com/Mayurhole95/LBMS/db"

type updateRequest struct {
	ID         string `json:"id"`
	First_Name string `json:"first_name"`
	Last_name  string `json:"last_name"`
	Gender     string `json:"gender"`
	Address    string `json:"address"`
	Password   string `json:"password"`
	Mob_no     string `json:"mob_no"`
}

type createRequest struct {
	ID         string `json:"id"`
	First_name string `json:"first_name"`
	Last_name  string `json:"last_name"`
	Gender     string `json:"gender"`
	Address    string `json:"address"`
	Email      string `json:"email"`
	Password   string `json:"password"`
	Mob_no     string `json:"mob_no"`
	Role       string `json:"role"`
}

type findByIDResponse struct {
	User db.User `json:"user"`
}

type listResponse struct {
	Users []db.User `json:"users"`
}

func (cr createRequest) Validate() (err error) {
	if cr.First_name == "" {
		return errEmptyName
	}
	return
}

func (ur updateRequest) Validate() (err error) {
	if ur.ID == "" {
		return errEmptyID
	}
	if ur.First_Name == "" {
		return errEmptyName
	}
	return
}
