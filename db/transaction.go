package db

import (
	"context"
	"database/sql"
	"fmt"
	"time"
)

const (
	createTransactionQuery = `INSERT INTO transaction ( 
    id,issuedate,duedate,returndate,book_id,user_id)
    VALUES(?,?,?,?,?,?)`

	listTransactionsQuery  = `SELECT * FROM transaction`
	UpdateTransactionQuery = `UPDATE transaction SET returndate=? WHERE book_id=? AND user_id=? AND returndate=0`
	issueCopyQuery         = `UPDATE book SET available_copies=available_copies-1 WHERE id = ? AND available_copies>0`
	returnCopyQuery        = `UPDATE book SET available_copies=available_copies+1 WHERE id = ?`
	BookStatusQuery        = `SELECT returndate from transaction WHERE book_id = ? AND user_id =?`
	GetTotalCopiesQuery    = `SELECT book.total_copies FROM book LEFT JOIN transaction ON book.id =transaction.book_id`
	GetCurrentCopiesQuery  = `SELECT book.available_copies FROM book LEFT JOIN transaction ON book.id =transaction.book_id where book.id=?`
)

type Transaction struct {
	ID         string `db:"id"`
	IssueDate  string `db:"issuedate"`
	DueDate    string `db:"duedate"`
	ReturnDate string `db:"returndate"`
	BookID     string `db:"book_id"`
	UserID     string `db:"user_id"`
}

func (s *store) CreateTransaction(ctx context.Context, transaction *Transaction) (err error) {

	// now := time.Now().UTC().Unix()
	// transaction.DueDate = int(now) + 864000
	now := time.Now().UTC().Unix()
	t := time.Unix(now, 0)
	strnow := t.Format(time.UnixDate)
	transactionduedate := int(now) + 864000
	t = time.Unix(int64(transactionduedate), 0)
	strduedate := t.Format(time.UnixDate)

	return Transact(ctx, s.db, &sql.TxOptions{}, func(ctx context.Context) error {
		_, err = s.db.Exec(
			createTransactionQuery,
			transaction.ID,
			strnow,
			strduedate,
			0,
			transaction.BookID,
			transaction.UserID,
		)
		if err != nil {
			return err
		}

		_, err = s.db.Exec(
			issueCopyQuery,
			transaction.BookID,
		)
		return err
	})
}

func (s *store) ListTransaction(ctx context.Context) (transactions []Transaction, err error) {
	err = WithDefaultTimeout(ctx, func(ctx context.Context) error {
		return s.db.SelectContext(ctx, &transactions, listTransactionsQuery)
	})
	if err == sql.ErrNoRows {
		return transactions, ErrTransactionNotExist
	}
	return
}

func (s *store) UpdateTransaction(ctx context.Context, transaction *Transaction) (err error) {
	//now := time.Now().UTC().Unix()
	now := time.Now().UTC().Unix()
	t := time.Unix(now, 0)
	strnow := t.Format(time.UnixDate)
	return Transact(ctx, s.db, &sql.TxOptions{}, func(ctx context.Context) error {
		_, err = s.db.Exec(
			UpdateTransactionQuery,

			strnow,
			transaction.BookID,
			transaction.UserID,
		)
		if err != nil {
			return err
		}

		_, err = s.db.Exec(
			returnCopyQuery,
			transaction.BookID,
		)
		return err
	})
}
func (s *store) BookStatus(ctx context.Context, BookId string, UserID string) (res string, err error) {
	return_date := -1
	s.db.GetContext(ctx, &return_date, BookStatusQuery, BookId, UserID)
	fmt.Println(return_date)
	if return_date == 0 {
		res = "issued"
		return res, nil
	} else {
		totalcnt := 0
		currentcnt := 0
		s.db.GetContext(ctx, &totalcnt, GetTotalCopiesQuery, BookId)
		s.db.GetContext(ctx, &currentcnt, GetCurrentCopiesQuery, BookId)

		if currentcnt < 1 {
			res = "Unavailable"
			return res, nil
		} else {
			res = "Available"
			return res, nil
		}
	}
}
