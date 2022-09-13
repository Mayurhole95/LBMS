package db

import (
	"context"
	"database/sql"
	"time"
)

const (
	createBookQuery = `INSERT INTO user ( 
    ID,Name,Author,Price,TotalCopies,Status,AvailableCopies)
    VALUES(?,?,?,?,?,?,?,?,?)`

	listBooksQuery      = `SELECT * FROM book`
	findBookByIDQuery   = `SELECT * FROM book WHERE id = ?`
	deleteBookByIDQuery = `DELETE FROM book WHERE id = ?`
	updateBookQuery     = `UPDATE book SET name = $1, updated_at = $2 where id = $3`
)

type Book struct {
	ID              string `json:"id"`
	Name            string `json:"name"`
	Author          string `json:"author"`
	Price           int    `json: "price"`
	TotalCopies     int    `json: "total_copies"`
	Status          string `json:"status"`
	AvailableCopies int    `json: "available_copies"`
}

func (s *store) CreateBOok(ctx context.Context, book *Book) (err error) {
	//now := time.Now()

	return Transact(ctx, s.db, &sql.TxOptions{}, func(ctx context.Context) error {
		_, err = s.db.Exec(
			createUserQuery,
			book.ID,
			book.Name,
			book.Author,
			book.Price,
			book.TotalCopies,
			book.Status,
			book.AvailableCopies,
		)
		return err
	})
}

func (s *store) ListBooks(ctx context.Context) (books []Book, err error) {
	err = WithDefaultTimeout(ctx, func(ctx context.Context) error {
		return s.db.SelectContext(ctx, &books, listBooksQuery)
	})
	if err == sql.ErrNoRows {
		return books, ErrBookNotExist
	}
	return
}

func (s *store) FindBookByID(ctx context.Context, id string) (book Book, err error) {
	err = WithDefaultTimeout(ctx, func(ctx context.Context) error {
		return s.db.GetContext(ctx, &book, findBookByIDQuery, id)
	})
	if err == sql.ErrNoRows {
		return book, ErrBookNotExist
	}
	return
}

func (s *store) DeleteBookByID(ctx context.Context, id string) (err error) {
	return Transact(ctx, s.db, &sql.TxOptions{}, func(ctx context.Context) error {
		res, err := s.db.Exec(deleteBookByIDQuery, id)
		cnt, err := res.RowsAffected()
		if cnt == 0 {
			return ErrBookNotExist
		}
		if err != nil {
			return err
		}
		return err
	})
}

func (s *store) UpdateBook(ctx context.Context, book *Book) (err error) {
	now := time.Now()

	return Transact(ctx, s.db, &sql.TxOptions{}, func(ctx context.Context) error {
		_, err = s.db.Exec(
			updateUserQuery,
			book.Name,
			now,
			book.ID,
		)
		return err
	})
}
