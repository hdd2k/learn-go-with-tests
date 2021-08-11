package main

import (
	"testing"
)

func TestWallet(t *testing.T) {

	assertBalance := func(t testing.TB, wallet Wallet, expected Bitcoin) {
		t.Helper()
		got := wallet.Balance()
		if got != expected {
			t.Errorf("got %s expected %s", got, expected)
		}
	}

	t.Run("Deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Bitcoin(10))
		assertBalance(t, wallet, Bitcoin(10))
	})

	t.Run("Withdraw", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(10)}
		wallet.Withdraw(Bitcoin(10))
		assertBalance(t, wallet, Bitcoin(0))
	})

	t.Run("Withdraw insufficient funds", func(t *testing.T) {
		startBal := Bitcoin(20)
		wallet := Wallet{startBal}
		err := wallet.Withdraw(Bitcoin(100))
		assertBalance(t, wallet, startBal)

		if err == nil {
			t.Error("Wanted err but didn't get one")
		}

	})
}
