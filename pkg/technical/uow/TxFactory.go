package uow

import (
	"context"
	"database/sql/driver"
)

type TxFactory interface {
	Tx() (driver.Tx, error)
	ContextTx(ctx context.Context) (ctxTx context.Context, tx driver.Tx, err error)
}

func HandleErrorByRollback(err error, tx driver.Tx) (rollbackErr error) {
	if tx == nil {
		return nil
	}

	if err != nil {
		if rollbackErr = tx.Rollback(); rollbackErr != nil {
			return rollbackErr
		}
	}
	return nil
}
