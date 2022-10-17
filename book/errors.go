package book

import "errors"

var (
	errEmptyID                 = errors.New("Book ID must be present")
	errEmptyName               = errors.New("Book name must be present")
	errNoBooks                 = errors.New("No books present")
	errNoBookId                = errors.New("Book is not present")
	errInvalidStatus           = errors.New("Invalid Status of Book")
	errEmptyAuthor             = errors.New("EMpty author")
	errZeroCopies              = errors.New("Zero Copies")
	errInvalidTotalCopies      = errors.New("Invalid Total copies")
	errInvalidPrice            = errors.New("Invalid Price")
	errInvalidAvailableCopies  = errors.New("INvalid Available Copies")
	err1InvalidAvailableCopies = errors.New("Invalid Available COpies")
)
