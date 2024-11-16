package mysql

import (
	"database/sql"
	"errors"
	"fmt"
	"log/slog"

	"github.com/charmingruby/txgo/internal/giftshop/core/logic"
)

func NewTransactionConsistencyProvider(db *sql.DB) *TransactionConsistencyProvider {
	return &TransactionConsistencyProvider{
		db: db,
	}
}

type TransactionConsistencyProvider struct {
	db *sql.DB
}

func (p *TransactionConsistencyProvider) Transact(txFunc func(params logic.TransactionalConsistencyParams) error) error {
	err := p.runInTx(p.db, func(tx *sql.Tx) error {
		return txFunc(logic.TransactionalConsistencyParams{
			GiftRepository:        NewGiftRepository(tx),
			PaymentRepository:     NewPaymentRepository(tx),
			TransactionRepository: NewTransactionRepository(tx),
			WalletRepository:      NewWalletRepository(tx),
		})
	})

	if err != nil {
		slog.Error(fmt.Sprintf("Transact function returned error: %s", err.Error()))
	}
	return err
}

func (p *TransactionConsistencyProvider) runInTx(db *sql.DB, fn func(*sql.Tx) error) error {
	tx, err := db.Begin()
	if err != nil {
		slog.Error(fmt.Sprintf("Failed to begin transaction: %s", err.Error()))
		return err
	}

	err = fn(tx)
	if err == nil {
		if commitErr := tx.Commit(); commitErr != nil {
			slog.Error(fmt.Sprintf("Failed to commit transaction: %s", commitErr.Error()))
			return commitErr
		}
		return nil
	}

	slog.Error(fmt.Sprintf("Transaction function returned error: %s", err.Error()))
	rollbackErr := tx.Rollback()
	if rollbackErr != nil {
		slog.Error(fmt.Sprintf("Failed to rollback transaction: %s", rollbackErr.Error()))
		return errors.Join(err, rollbackErr)
	}

	return err
}
