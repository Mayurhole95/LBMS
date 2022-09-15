package transaction

import (
	"context"

	"github.com/Mayurhole95/LBMS/db"
	"github.com/google/uuid"
	"go.uber.org/zap"
)

type Service interface {
	list(ctx context.Context) (response listResponse, err error)
	create(ctx context.Context, req Transaction) (err error)
	delete(ctx context.Context, id string) (err error)
	update(ctx context.Context, req Transaction) (err error)
}

type transactionService struct {
	store  db.Storer
	logger *zap.SugaredLogger
}

func (cs *transactionService) list(ctx context.Context) (response listResponse, err error) {
	transactions, err := cs.store.ListTransaction(ctx)
	if err == db.ErrTransactionNotExist {
		cs.logger.Error("No transaction present", "err", err.Error())
		return response, errNoTransactions
	}
	if err != nil {
		cs.logger.Error("Error listing transactions", "err", err.Error())
		return
	}

	response.Transactions = transactions
	return
}

func (cs *transactionService) create(ctx context.Context, c Transaction) (err error) {
	err = c.Validate()
	if err != nil {
		cs.logger.Errorw("Invalid request for transaction creation", "msg", err.Error(), "transaction", c)
		return
	}
	uuidgen := uuid.New()
	c.ID = uuidgen.String()
	err = cs.store.CreateTransaction(ctx, &db.Transaction{
		ID:         c.ID,
		IssueDate:  c.IssueDate,
		DueDate:    c.DueDate,
		ReturnDate: c.ReturnDate,
		BookID:     c.BookID,
		UserID:     c.UserID,
	})
	if err != nil {
		cs.logger.Error("Error creating transaction", "err", err.Error())
		return
	}
	return
}

func (cs *transactionService) update(ctx context.Context, c Transaction) (err error) {

	err = cs.store.UpdateTransaction(ctx, &db.Transaction{

		ReturnDate: c.ReturnDate,
		BookID:     c.BookID,
		UserID:     c.UserID,
	})
	if err != nil {
		cs.logger.Error("Error updating transaction", "err", err.Error(), "transaction", c)
		return
	}

	return
}

func (cs *transactionService) delete(ctx context.Context, id string) (err error) {
	err = cs.store.DeleteTransaction(ctx, id)
	if err == db.ErrTransactionNotExist {
		cs.logger.Error("Transaction Not present", "err", err.Error(), "transaction_id", id)
		return errNoTransactionId
	}
	if err != nil {
		cs.logger.Error("Error deleting transaction", "err", err.Error(), "transaction_id", id)
		return
	}

	return
}

func NewService(s db.Storer, l *zap.SugaredLogger) Service {
	return &transactionService{
		store:  s,
		logger: l,
	}
}
