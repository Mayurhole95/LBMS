package user

import "errors"

var (
	errEmptyID          = errors.New("User ID must be present")
	errEmptyName        = errors.New("User name must be present")
	errNoUsers          = errors.New("No users present")
	errNoUserId         = errors.New("User is not present")
	errEmptyPassword    = errors.New("Password cannot be empty")
	errInvalidGender    = errors.New("Enter valid gender")
	errEmptyAddress     = errors.New("Address must be present")
	errEmptyEmail       = errors.New("Email must be present")
	errEmptyMobNo       = errors.New("Mob no must be present")
	errEmptyRole        = errors.New("Role must be present")
	errRoleType         = errors.New("Enter a valid Role type")
	errNotValidMail     = errors.New("Email is not valid")
	errInvalidMobNo     = errors.New("Mob Number is not valid")
	errEmptyLastName    = errors.New("Last Name cannot be empty")
	errBookALreadyTaken = errors.New("Book taken so cannot delete user")
	errWrongPassword    = errors.New("Wrong Password")
	errInvalidFirstName = errors.New("Invalid First Name")
	errInvalidLastName  = errors.New("Invalid Last Name")
	InvalidPassword     = errors.New("Password should be greater than 6 characters")
)
