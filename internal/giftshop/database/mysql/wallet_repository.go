package mysql

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/charmingruby/txgo/internal/giftshop/core/model"
	"github.com/charmingruby/txgo/internal/shared/core/core_err"
)

const (
	WALLETS_TABLE = "wallets"
)

func NewWalletRepository(db *sql.DB) *WalletRepository {
	return &WalletRepository{db: db}
}

type WalletRepository struct {
	db *sql.DB
}

type walletRow struct {
	ID         string
	Name       string
	OwnerEmail string
	Points     int
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

func (r *WalletRepository) FindByOwnerEmail(ownerEmail string) (*model.Wallet, error) {
	query := fmt.Sprintf("SELECT * FROM %s WHERE owner_email = ?", WALLETS_TABLE)

	queryResult := r.db.QueryRow(query, ownerEmail)

	var row walletRow

	if err := queryResult.Scan(
		&row.ID,
		&row.Name,
		&row.OwnerEmail,
		&row.Points,
		&row.CreatedAt,
		&row.UpdatedAt,
	); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}

		return nil, core_err.NewPersistenceErr(err, "wallet find_by_owner_email", "mysql")
	}

	return r.mapToDomain(row), nil
}

func (r *WalletRepository) Store(wallet *model.Wallet) error {
	query := fmt.Sprintf("INSERT INTO %s (id, name, owner_email, points, created_at, updated_at) VALUES (?, ?, ?, ?, ?, ?)", WALLETS_TABLE)

	_, err := r.db.Exec(query, wallet.ID(), wallet.Name(), wallet.OwnerEmail(), wallet.Points(), wallet.CreatedAt(), wallet.UpdatedAt())
	if err != nil {
		return core_err.NewPersistenceErr(err, "wallet store", "mysql")
	}

	return nil
}

func (r *WalletRepository) mapToDomain(wallet walletRow) *model.Wallet {
	input := model.NewWalletFromInput{
		ID:         wallet.ID,
		CreatedAt:  wallet.CreatedAt,
		UpdatedAt:  wallet.UpdatedAt,
		Name:       wallet.Name,
		OwnerEmail: wallet.OwnerEmail,
		Points:     wallet.Points,
	}

	return model.NewWalletFrom(input)
}
