package factory

import (
	"time"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/charmingruby/txgo/internal/giftshop/core/model"
	"github.com/charmingruby/txgo/internal/giftshop/core/repository"
	"github.com/charmingruby/txgo/internal/shared/core"
	"github.com/charmingruby/txgo/internal/shared/helper"
)

func MakeWallet(walletRepo repository.WalletRepository, params model.NewWalletFromInput) (*model.Wallet, error) {
	wallet := createWallet(params)

	if err := walletRepo.Store(wallet); err != nil {
		return nil, err
	}

	return wallet, nil
}

func createWallet(params model.NewWalletFromInput) *model.Wallet {
	input := model.NewWalletFromInput{
		ID:         helper.If[string](params.ID != "", params.ID, core.NewID()),
		Name:       helper.If[string](params.Name != "", params.Name, gofakeit.Name()),
		OwnerEmail: helper.If[string](params.OwnerEmail != "", params.OwnerEmail, gofakeit.Email()),
		Points:     params.Points,
		CreatedAt:  helper.If[time.Time](params.CreatedAt != time.Time{}, params.CreatedAt, time.Now()),
		UpdatedAt:  helper.If[time.Time](params.UpdatedAt != time.Time{}, params.UpdatedAt, time.Now()),
	}

	return model.NewWalletFrom(input)
}
