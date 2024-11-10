package mysql

import (
	"database/sql"
	"fmt"

	"github.com/charmingruby/txgo/internal/giftshop/core/model"
	"github.com/charmingruby/txgo/internal/shared/core/core_err"
)

const (
	PAYMENTS_TABLE = "payments"
)

func NewPaymentRepository(db *sql.DB) *PaymentRepository {
	return &PaymentRepository{db: db}
}

type PaymentRepository struct {
	db *sql.DB
}

func (r *PaymentRepository) Store(payment *model.Payment) error {
	query := fmt.Sprintf("INSERT INTO %s (id, installments, tax_percent, partial_value, total_value, status, transaction_id, created_at, updated_at) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)", PAYMENTS_TABLE)

	transactionID := sql.NullString{String: payment.TransactionID(), Valid: payment.TransactionID() != ""}

	_, err := r.db.Exec(query,
		payment.ID(),
		payment.Installments(),
		payment.TaxPercent(),
		payment.PartialValue(),
		payment.TotalValue(),
		payment.Status(),
		transactionID,
		payment.CreatedAt(),
		payment.UpdatedAt())
	if err != nil {
		return core_err.NewPersistenceErr(err, "payment store", "mysql")
	}

	return nil
}

func (r *PaymentRepository) UpdateTransactionIDAndStatusByID(payment *model.Payment) error {
	query := fmt.Sprintf("UPDATE %s SET transaction_id = ?, status = ?, updated_at = ? WHERE id = ?", PAYMENTS_TABLE)

	_, err := r.db.Exec(query, payment.TransactionID(), payment.Status(), payment.UpdatedAt(), payment.ID())
	if err != nil {
		return core_err.NewPersistenceErr(err, "payment update_transaction_id_and_status_by_id", "mysql")
	}

	return nil
}
