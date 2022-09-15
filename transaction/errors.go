package transaction

import "errors"

var (
	errEmptyID         = errors.New("ID must be present")
	errNoTransactions  = errors.New("No transaction present")
	errNoTransactionId = errors.New("Transaction ID is not present")
)
