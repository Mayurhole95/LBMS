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

// func (s *store) DeleteTransaction(ctx context.Context, id string) (err error) {
// 	return Transact(ctx, s.db, &sql.TxOptions{}, func(ctx context.Context) error {
// 		res, err := s.db.Exec(DeleteTransactionQuery, id)
// 		cnt, err := res.RowsAffected()
// 		if cnt == 0 {
// 			return ErrTransactionNotExist
// 		}
// 		if err != nil {
// 			return err
// 		}
// 		return err
// 	})
// }

func (s *store) UpdateTransaction(ctx context.Context, transaction *Transaction) (err error) {

	return Transact(ctx, s.db, &sql.TxOptions{}, func(ctx context.Context) error {
		_, err = s.db.Exec(
			UpdateTransactionQuery,
			transaction.ReturnDate,
			transaction.BookID,
			transaction.UserID,
		)
		return err
	})
}
