package transaction

import "github.com/Mayurhole95/LBMS/db"

type Transaction struct {
	ID         string `json:"id"`
	IssueDate  int    `json:"issuedate"`
	DueDate    int    `json:"duedate"`
	ReturnDate int    `json:"returndate"`
	BookID     string `json:"book_id"`
	UserID     string `json:"user_id"`
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
