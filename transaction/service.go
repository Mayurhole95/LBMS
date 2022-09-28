package transaction

import (
	"context"

	"github.com/Mayurhole95/LBMS/db"
	"github.com/google/uuid"
	"go.uber.org/zap"
)

type Service interface {
	List(ctx context.Context) (response listResponse, err error)
	Create(ctx context.Context, req Transaction) (err error)
	Update(ctx context.Context, req Transaction) (err error)
	BookStatus(ctx context.Context, c RequestStatus) (response string, err error)
}

type transactionService struct {
	store  db.Storer
	logger *zap.SugaredLogger
}

func (cs *transactionService) List(ctx context.Context) (response listResponse, err error) {
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

func (cs *transactionService) Create(ctx context.Context, c Transaction) (err error) {
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

func (cs *transactionService) Update(ctx context.Context, c Transaction) (err error) {

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

func (cs *transactionService) BookStatus(ctx context.Context, c RequestStatus) (response string, err error) {
	response, err = cs.store.BookStatus(ctx, c.BookID, c.UserID)
	if err == db.ErrUserNotExist {
		cs.logger.Error("No Transaction present", "err", err.Error())
		return response, errNoTransaction
	}
	if err != nil {
		cs.logger.Error("Error listing Transactions", "err", err.Error())
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
