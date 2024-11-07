package service

import "github.com/charmingruby/txgo/internal/giftshop/core/repository"

type Service struct {
	//paymentRepo     repository.PaymentRepository
	//giftRepo        repository.GiftRepository
	walletRepo repository.WalletRepository
	//transactionRepo repository.TransactionRepository
}

func NewService(
	//paymentRepo repository.PaymentRepository,
	//giftRepo repository.GiftRepository,
	walletRepo repository.WalletRepository,
	//transactionRepo repository.TransactionRepository,
) *Service {
	return &Service{
		//paymentRepo:     paymentRepo,
		//giftRepo:        giftRepo,
		walletRepo: walletRepo,
		//transactionRepo: transactionRepo,
	}
}
