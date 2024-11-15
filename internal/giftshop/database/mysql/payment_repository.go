package mysql

import (
	"database/sql"
	"fmt"
	"time"

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

type paymentRow struct {
	ID            string
	Installments  int
	TaxPercent    int
	PartialValue  int
	TotalValue    int
	Status        string
	TransactionID string
	CreatedAt     time.Time
	UpdatedAt     time.Time
}

func (r *PaymentRepository) FindByID(id string) (*model.Payment, error) {
	query := fmt.Sprintf("SELECT * FROM %s WHERE id = ?", PAYMENTS_TABLE)

	queryResult := r.db.QueryRow(query, id)

	var row paymentRow

	if err := queryResult.Scan(
		&row.ID,
		&row.Installments,
		&row.TaxPercent,
		&row.PartialValue,
		&row.TotalValue,
		&row.Status,
		&row.TransactionID,
		&row.CreatedAt,
		&row.UpdatedAt,
	); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}

		return nil, core_err.NewPersistenceErr(err, "payment find_by_id", "mysql")
	}

	return r.mapToDomain(row), nil
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

func (r *PaymentRepository) mapToDomain(payment paymentRow) *model.Payment {
	return model.NewPaymentFrom(model.NewPaymentFromInput{
		ID:            payment.ID,
		Installments:  payment.Installments,
		TaxPercent:    payment.TaxPercent,
		PartialValue:  payment.PartialValue,
		Status:        payment.Status,
		TotalValue:    payment.TotalValue,
		TransactionID: payment.TransactionID,
		CreatedAt:     payment.CreatedAt,
		UpdatedAt:     payment.UpdatedAt,
	})
}
