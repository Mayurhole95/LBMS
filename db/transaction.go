package db

import (
	"context"
	"database/sql"
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
)

type Transaction struct {
	ID         string `db:"id"`
	IssueDate  int    `db:"issuedate"`
	DueDate    int    `db:"duedate"`
	ReturnDate int    `db:"returndate"`
	BookID     string `db:"book_id"`
	UserID     string `db:"user_id"`
}

func (s *store) CreateTransaction(ctx context.Context, transaction *Transaction) (err error) {

	now := time.Now().UTC().Unix()
	transaction.DueDate = int(now) + 864000
	return Transact(ctx, s.db, &sql.TxOptions{}, func(ctx context.Context) error {
		_, err = s.db.Exec(
			createTransactionQuery,
			transaction.ID,
			now,
			transaction.DueDate,
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
	now := time.Now().UTC().Unix()
	return Transact(ctx, s.db, &sql.TxOptions{}, func(ctx context.Context) error {
		_, err = s.db.Exec(
			UpdateTransactionQuery,

			now,
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
