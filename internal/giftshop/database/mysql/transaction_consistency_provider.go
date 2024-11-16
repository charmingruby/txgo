package mysql

import (
	"database/sql"
	"errors"

	"github.com/charmingruby/txgo/internal/giftshop/core/logic"
)

func NewTransactionConsistencyProvider(db *sql.DB) *TransactionConsistencyProvider {
	return &TransactionConsistencyProvider{
		db: db,
	}
}

type TransactionConsistencyProvider struct {
	db      *sql.DB
	context logic.TransactionalConsistencyLogicParams
}

func (p *TransactionConsistencyProvider) Transact(params logic.TransactionalConsistencyLogicParams) {
	return
}

func (p *TransactionConsistencyProvider) runInTx(db *sql.DB, fn func(*sql.Tx) error) error {
	tx, err := db.Begin()
	if err != nil {
		return err
	}

	if err := fn(tx); err == nil {
		return tx.Commit()
	}

	rollbackErr := tx.Rollback()
	if rollbackErr != nil {
		return errors.Join(err, rollbackErr)
	}

	return err
}
