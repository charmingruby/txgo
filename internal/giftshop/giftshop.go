package giftshop

import (
	"database/sql"

	"github.com/charmingruby/txgo/internal/giftshop/core/repository"
	"github.com/charmingruby/txgo/internal/giftshop/core/service"
	"github.com/charmingruby/txgo/internal/giftshop/database/mysql"
	"github.com/charmingruby/txgo/internal/giftshop/transport/rest/endpoint"
	"github.com/go-chi/chi/v5"
)

func NewService(
	walletRepository repository.WalletRepository,
	giftRepository repository.GiftRepository,
) *service.Service {
	return service.NewService(giftRepository, walletRepository)
}

func NewWalletRepository(db *sql.DB) repository.WalletRepository {
	return mysql.NewWalletRepository(db)
}

func NewGiftRepository(db *sql.DB) repository.GiftRepository {
	return mysql.NewGiftRepository(db)
}

func NewHTTPHandler(r *chi.Mux, service *service.Service) {
	endpoint.New(r, service).Register()
}
