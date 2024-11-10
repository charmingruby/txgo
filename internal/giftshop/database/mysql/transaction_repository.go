package mysql

import (
	"database/sql"
	"fmt"

	"github.com/charmingruby/txgo/internal/giftshop/core/model"
	"github.com/charmingruby/txgo/internal/shared/core/core_err"
)

const (
	TRANSACTIONS_TABLE = "transactions"
)

func NewTransactionRepository(db *sql.DB) *TransactionRepository {
	return &TransactionRepository{db: db}
}

type TransactionRepository struct {
	db *sql.DB
}

func (r *TransactionRepository) Store(transaction *model.Transaction) error {
	query := fmt.Sprintf("INSERT INTO %s (id, points, payer_wallet_id, created_at, updated_at) VALUES (?, ?, ?, ?, ?)", TRANSACTIONS_TABLE)

	_, err := r.db.Exec(query,
		transaction.ID(),
		transaction.Points(),
		transaction.PayerWalletID(),
		transaction.CreatedAt(),
		transaction.UpdatedAt(),
	)
	if err != nil {
		return core_err.NewPersistenceErr(err, "transaction store", "mysql")
	}

	return nil
}
