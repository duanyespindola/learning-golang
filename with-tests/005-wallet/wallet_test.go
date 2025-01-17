package wallet

import "testing"

var assertError = func(t testing.TB, err error) {
	t.Helper()
	if err == nil {
		t.Error("wanted an error but didn't get one")
	}
}

func TestWallet(t *testing.T) {

	assertBalance := func(t testing.TB, wallet Wallet, want Bitcoin) {
		t.Helper()
		got := wallet.Balance()

		if got != want {
			t.Errorf("got %s want %s", got, want)
		}
	}
	t.Run("Deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Bitcoin(10))
		assertBalance(t, wallet, Bitcoin(10))
	})

	t.Run("Withdraw", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(20)}
		err := wallet.Withdraw(Bitcoin(8))
		if err != nil {
			t.Errorf("There was a error with withdraw")
		} else {
			assertBalance(t, wallet, Bitcoin(12))
		}
	})

	t.Run("withdraw insufficient funds", func(t *testing.T) {
		startingBalance := Bitcoin(20)
		wallet := Wallet{startingBalance}
		err := wallet.Withdraw(Bitcoin(100))
		assertError(t, err)
		assertBalance(t, wallet, startingBalance)
	})

}
