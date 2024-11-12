package endpoint

import (
	"github.com/charmingruby/txgo/test/factory"
)

func (s *Suite) Test_CreateWalletHandler() {
	s.Run("it should be able to create a new wallet", func() {
		wallet, err := factory.MakeWallet(s.walletRepo, factory.MakeWalletParams{})
		s.NoError(err)

		w, err := s.walletRepo.FindByID(wallet.ID())
		s.NoError(err)
		s.NotNil(w)
	})
}
