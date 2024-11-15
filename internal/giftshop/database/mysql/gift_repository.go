package mysql

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/charmingruby/txgo/internal/giftshop/core/model"
	"github.com/charmingruby/txgo/internal/shared/core/core_err"
	"github.com/charmingruby/txgo/internal/shared/helper"
)

const (
	GIFTS_TABLE = "gifts"
)

func NewGiftRepository(db *sql.DB) *GiftRepository {
	return &GiftRepository{db: db}
}

type GiftRepository struct {
	db *sql.DB
}

type giftRow struct {
	ID               string
	Name             string
	Message          string
	ReceiverWalletID string
	SenderWalletID   string
	BaseValue        int
	Status           string
	PaymentID        string
	CreatedAt        time.Time
	UpdatedAt        time.Time
}

func (r *GiftRepository) FindByID(id string) (*model.Gift, error) {
	query := fmt.Sprintf("SELECT * FROM %s WHERE id = ?", GIFTS_TABLE)

	queryResult := r.db.QueryRow(query, id)

	var row giftRow

	var paymentID sql.NullString

	if err := queryResult.Scan(
		&row.ID,
		&row.Name,
		&row.Message,
		&row.BaseValue,
		&row.Status,
		&row.ReceiverWalletID,
		&row.SenderWalletID,
		&paymentID,
		&row.CreatedAt,
		&row.UpdatedAt,
	); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}

		return nil, core_err.NewPersistenceErr(err, "gift find_by_id", "mysql")
	}

	row.PaymentID = helper.If[string](paymentID.Valid, paymentID.String, "")

	return r.mapToDomain(row), nil
}

func (r *GiftRepository) Store(gift *model.Gift) error {
	query := fmt.Sprintf("INSERT INTO %s (id, name, message, base_value, status, receiver_wallet_id, sender_wallet_id, payment_id, created_at, updated_at) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)", GIFTS_TABLE)

	paymentID := sql.NullString{String: gift.PaymentID(), Valid: gift.PaymentID() != ""}

	_, err := r.db.Exec(query, gift.ID(), gift.Name(), gift.Message(), gift.BaseValue(), gift.Status(), gift.ReceiverWalletID(), gift.SenderWalletID(), paymentID, gift.CreatedAt(), gift.UpdatedAt())
	if err != nil {
		return core_err.NewPersistenceErr(err, "gift store", "mysql")
	}

	return nil
}

func (r *GiftRepository) UpdatePaymentIDAndStatusByID(gift *model.Gift) error {
	query := fmt.Sprintf("UPDATE %s SET payment_id = ?, status = ?, updated_at = ? WHERE id = ?", GIFTS_TABLE)

	_, err := r.db.Exec(query, gift.PaymentID(), gift.Status(), gift.UpdatedAt(), gift.ID())
	if err != nil {
		return core_err.NewPersistenceErr(err, "gift update_payment_id_and_status_by_id", "mysql")
	}

	return nil
}

func (r *GiftRepository) mapToDomain(gift giftRow) *model.Gift {
	return model.NewGiftFrom(model.NewGiftFromInput{
		ID:               gift.ID,
		Name:             gift.Name,
		Message:          gift.Message,
		BaseValue:        gift.BaseValue,
		Status:           gift.Status,
		PaymentID:        gift.PaymentID,
		SenderWalletID:   gift.SenderWalletID,
		ReceiverWalletID: gift.ReceiverWalletID,
		CreatedAt:        gift.CreatedAt,
		UpdatedAt:        gift.UpdatedAt,
	})
}
