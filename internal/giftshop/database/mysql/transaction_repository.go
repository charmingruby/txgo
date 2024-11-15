package mysql

import (
	"database/sql"
	"fmt"
	"time"

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

type transactionRow struct {
	ID            string
	Points        int
	PayerWalletID string
	CreatedAt     time.Time
	UpdatedAt     time.Time
}

func (r *TransactionRepository) FindByID(id string) (*model.Transaction, error) {
	query := fmt.Sprintf("SELECT * FROM %s WHERE id = ?", TRANSACTIONS_TABLE)

	queryResult := r.db.QueryRow(query, id)

	var row transactionRow

	if err := queryResult.Scan(
		&row.ID,
		&row.Points,
		&row.PayerWalletID,
		&row.CreatedAt,
		&row.UpdatedAt,
	); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}

		return nil, core_err.NewPersistenceErr(err, "transaction find_by_id", "mysql")
	}

	return r.mapToDomain(row), nil
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

func (r *TransactionRepository) mapToDomain(transaction transactionRow) *model.Transaction {
	return model.NewTransactionFrom(model.NewTransactionFromInput{
		ID:            transaction.ID,
		Points:        transaction.Points,
		PayerWalletID: transaction.PayerWalletID,
		CreatedAt:     transaction.CreatedAt,
		UpdatedAt:     transaction.UpdatedAt,
	})
}
