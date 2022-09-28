package user

import (
	"net/mail"
	"unicode"

	"github.com/Mayurhole95/LBMS/db"
)

type UpdateRequest struct {
	ID         string `json:"id"`
	First_Name string `json:"first_name"`
	Last_name  string `json:"last_name"`
	Gender     string `json:"gender"`
	Address    string `json:"address"`
	Password   string `json:"password"`
	Mob_no     string `json:"mob_no"`
}
type ResetRequest struct {
	ID          string `json:"id"`
	Password    string `json:"password"`
	NewPassword string `json:"newpassword"`
}

type CreateRequest struct {
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

type FindByIDResponse struct {
	User db.User `json:"user"`
}

type ListResponse struct {
	Users []db.User `json:"users"`
}

func (cr CreateRequest) Validate() (err error) {
	if cr.First_name == "" {
		return errEmptyName
	}
	for _, r := range cr.First_name {
		if !unicode.IsLetter(r) {
			return errInvalidFirstName
		}
	}
	if cr.Last_name == "" {
		return errEmptyLastName
	}
	for _, r := range cr.Last_name {
		if !unicode.IsLetter(r) {
			return errInvalidLastName
		}
	}
	if cr.Password == "" {
		return errEmptyPassword
	}
	if len(cr.Password) < 6 {
		return InvalidPassword
	}

	if cr.Gender == "" || cr.Gender != "Male" && cr.Gender != "male" && cr.Gender != "Female" && cr.Gender != "female" && cr.Gender != "other" && cr.Gender != "Other" {
		return errInvalidGender
	}
	if cr.Address == "" {
		return errEmptyAddress
	}
	if cr.Email == "" {
		return errEmptyEmail
	}
	if cr.Mob_no == "" {
		return errEmptyMobNo
	}
	if cr.Role == "" {
		return errEmptyRole
	}
	if cr.Role != "user" && cr.Role != "admin" {
		return errRoleType
	}
	_, b := mail.ParseAddress(cr.Email)
	if b != nil {
		return errNotValidMail
	}
	checkEmail := cr.Email
	flag := false
	lastapperance := 0
	for i := 0; i < len(checkEmail); i++ {
		if checkEmail[i] == '@' {
			flag = true
			lastapperance = i
		}
	}
	if !flag {
		return errNotValidMail
	}
	flag = false
	for i := lastapperance; i < len(checkEmail); i++ {
		if checkEmail[i] == '.' {
			flag = true
		}
	}
	if !flag {
		return errNotValidMail
	}
	if len(cr.Mob_no) < 10 || len(cr.Mob_no) > 10 {
		return errInvalidMobNo
	}
	return
}

func (ur UpdateRequest) Validate() (err error) {
	if ur.ID == "" {
		return errEmptyID
	}
	if ur.First_Name == "" {
		return errEmptyName
	}

	return
}
