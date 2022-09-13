package book

import (
	"context"

	"github.com/Mayurhole95/LBMS/db"

	"go.uber.org/zap"
)

type Service interface {
	list(ctx context.Context) (response listResponse, err error)
	create(ctx context.Context, req createRequest) (err error)
	findByID(ctx context.Context, id string) (response findByIDResponse, err error)
	deleteByID(ctx context.Context, id string) (err error)
	update(ctx context.Context, req updateRequest) (err error)
}

type bookService struct {
	store  db.Storer
	logger *zap.SugaredLogger
}

func (cs *bookService) list(ctx context.Context) (response listResponse, err error) {
	books, err := cs.store.ListBooks(ctx)
	if err == db.ErrBookNotExist {
		cs.logger.Error("No book present", "err", err.Error())
		return response, errNoBooks
	}
	if err != nil {
		cs.logger.Error("Error listing books", "err", err.Error())
		return
	}

	response.Books = books
	return
}

func (cs *bookService) create(ctx context.Context, c createRequest) (err error) {
	err = c.Validate()
	if err != nil {
		cs.logger.Errorw("Invalid request for book creation", "msg", err.Error(), "book", c)
		return
	}

	err = cs.store.CreateBook(ctx, &db.Book{
		ID:              c.ID,
		Name:            c.Name,
		Author:          c.Author,
		Price:           c.Price,
		TotalCopies:     c.TotalCopies,
		Status:          c.Status,
		AvailableCopies: c.AvailableCopies,
	})
	if err != nil {
		cs.logger.Error("Error creating book", "err", err.Error())
		return
	}
	return
}

func (cs *bookService) update(ctx context.Context, c updateRequest) (err error) {
	err = c.Validate()
	if err != nil {
		cs.logger.Error("Invalid Request for book update", "err", err.Error(), "book", c)
		return
	}

	err = cs.store.UpdateBook(ctx, &db.Book{
		ID:              c.ID,
		Name:            c.Name,
		Author:          c.Author,
		Price:           c.Price,
		TotalCopies:     c.TotalCopies,
		Status:          c.Status,
		AvailableCopies: c.AvailableCopies,
	})
	if err != nil {
		cs.logger.Error("Error updating book", "err", err.Error(), "book", c)
		return
	}

	return
}

func (cs *bookService) findByID(ctx context.Context, id string) (response findByIDResponse, err error) {
	book, err := cs.store.FindBookByID(ctx, id)
	if err == db.ErrBookNotExist {
		cs.logger.Error("No user present", "err", err.Error())
		return response, errNoBookId
	}
	if err != nil {
		cs.logger.Error("Error finding book", "err", err.Error(), "book_id", id)
		return
	}

	response.Book = book
	return
}

func (cs *bookService) deleteByID(ctx context.Context, id string) (err error) {
	err = cs.store.DeleteBookByID(ctx, id)
	if err == db.ErrBookNotExist {
		cs.logger.Error("User Not present", "err", err.Error(), "book_id", id)
		return errNoBookId
	}
	if err != nil {
		cs.logger.Error("Error deleting book", "err", err.Error(), "book_id", id)
		return
	}

	return
}

func NewService(s db.Storer, l *zap.SugaredLogger) Service {
	return &bookService{
		store:  s,
		logger: l,
	}
}
