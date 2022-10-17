package transaction

import "github.com/Mayurhole95/LBMS/db"

type Transaction struct {
	ID         string `json:"id"`
	IssueDate  string `json:"issuedate"`
	DueDate    string `json:"duedate"`
	ReturnDate string `json:"returndate"`
	BookID     string `json:"book_id"`
	UserID     string `json:"user_id"`
}
type RequestStatus struct {
	UserID string `json:"user_id"`
	BookID string `json:"book_id"`
}

type listResponse struct {
	Transactions []db.Transaction `json:"transactions"`
}

func (cr Transaction) Validate() (err error) {
	if cr.ID == "" {
		return errEmptyID
	}
	return
}
