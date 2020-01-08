package ptrAndErr

import (
	"testing"
)

func assertCorrectionMsg(t2 *testing.T, wallet Wallet, want Bitcoin) {
	t2.Helper()
	got := wallet.Balance()
	if got != want {
		t2.Errorf("got %s want %s", got, want)
	}
}

func assertErr(t *testing.T, err error, want error) {
	t.Helper()
	if err == nil {
		t.Errorf("wanted an error but didnt get one")
	}

	if err.Error() != want.Error() {
		t.Errorf("want %s ,got %s", want.Error(), err.Error())
	}
}

func assertNoErr(t *testing.T, got error) {
	if got != nil {
		t.Errorf("got an error but didnt want one")
	}
}

func TestWallet(t *testing.T) {

	t.Run("Deposit", func(t *testing.T) {
		wallet := Wallet{}

		wallet.Deposit(Bitcoin(10))

		assertCorrectionMsg(t, wallet, Bitcoin(10))
	})

	t.Run("Withdraw", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(20)}

		got := wallet.Withdraw(Bitcoin(10))

		assertNoErr(t, got)

		assertCorrectionMsg(t, wallet, Bitcoin(10))
	})

	t.Run("withdraw insufficient funds", func(t *testing.T) {
		startingBanlace := Bitcoin(20)
		wallet := Wallet{startingBanlace}
		err := wallet.Withdraw(Bitcoin(100))

		assertCorrectionMsg(t, wallet, startingBanlace)

		assertErr(t, err, insufficientErr)

	})
}
