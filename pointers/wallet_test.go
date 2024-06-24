package main

import (
	"testing"
)

func TestWallet(t *testing.T) {
	assertError := func (t testing.TB, got, want error) {
		t.Helper()
		if got == nil {
			t.Fatal("Wanted an error but did not get one")
		}
		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	}
	assertBalance := func (t testing.TB, wallet Wallet, want Bitcoin) {
		t.Helper()
		got := wallet.Balance()
		if got != want {
			t.Errorf("got %s want %s", got, want)
		}
	}

	t.Run("deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Bitcoin(10))
		assertBalance(t, wallet, Bitcoin(10))
	})

	t.Run("withdraw", func(t *testing.T) {
		wallet := Wallet{balance: 100}
		wallet.Withdraw(Bitcoin(50))
		assertBalance(t, wallet, Bitcoin(50))
	})

	t.Run("withdraw insufficient funds", func(t *testing.T) {
		startingBalance := Bitcoin(15)
		wallet := Wallet{balance: startingBalance}
		err := wallet.Withdraw(Bitcoin(20))

		assertError(t, err, ErrInsuficientFunds)
		assertBalance(t, wallet, startingBalance)
	})
}