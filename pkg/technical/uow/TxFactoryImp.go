package uow

import (
	"context"
	"database/sql/driver"

	"github.com/jmoiron/sqlx"
)

func NewTxFactory(db *sqlx.DB) *TxFactoryImp {
	return &TxFactoryImp{db: db}
}

type TxFactoryImp struct {
	db *sqlx.DB
}

func (f *TxFactoryImp) Tx() (driver.Tx, error) {
	return f.db.Beginx()
}

func (f *TxFactoryImp) ContextTx(ctx context.Context) (ctxTx context.Context, tx driver.Tx, err error) {
	tx, err = f.db.Beginx()
	if err != nil {
		return ctx, nil, err
	}
	ctxTx = context.WithValue(ctx, f.db, tx) // 表示 tx 連線, 來自同一個 db
	return ctxTx, tx, nil
}

func GetDBOrTx(db *sqlx.DB, tx driver.Tx) (RDBMS sqlx.ExtContext) {
	if tx == nil {
		return db
	}
	externalTx := tx.(*sqlx.Tx)
	return externalTx
}

func GetDBOrTxViaContext(db *sqlx.DB, ctx context.Context) (RDBMS sqlx.ExtContext) {
	if ctx == nil {
		panic("ctx context.Context is nil")
	}

	externalTx, ok := ctx.Value(db).(*sqlx.Tx)
	if !ok {
		return db
	}
	return externalTx
}
