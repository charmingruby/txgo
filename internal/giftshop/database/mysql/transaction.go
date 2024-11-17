package mysql

import (
	"database/sql"

	"github.com/charmingruby/txgo/internal/giftshop/core/service"
	"github.com/charmingruby/txgo/internal/shared/database/mysql"
)

func NewTransactionConsistencyProvider(db *sql.DB) *TransactionConsistencyProvider {
	return &TransactionConsistencyProvider{
		db: db,
	}
}

type TransactionConsistencyProvider struct {
	db *sql.DB
}

func (p *TransactionConsistencyProvider) Transact(txFunc func(params service.TransactionalConsistencyParams) error) error {
	err := mysql.RunInTx(p.db, func(tx *sql.Tx) error {
		return txFunc(service.TransactionalConsistencyParams{
			GiftRepository:        NewGiftRepository(tx),
			PaymentRepository:     NewPaymentRepository(tx),
			TransactionRepository: NewTransactionRepository(tx),
			WalletRepository:      NewWalletRepository(tx),
		})
	})

	return err
}
