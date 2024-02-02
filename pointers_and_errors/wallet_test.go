package pointers_and_errors

import (
	"errors"
	"testing"
)

func TestWallet(t *testing.T) {

	t.Run("deposit", func(t *testing.T) {
		wallet := Wallet{}

		wallet.Deposit(Bitcoin(10))

		expected := Bitcoin(10)
		assertBalance(t, wallet, expected)
	})

	t.Run("withdraw", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Bitcoin(20))

		err := wallet.Withdraw(Bitcoin(10))

		expected := Bitcoin(10)
		assertNoError(t, err)
		assertBalance(t, wallet, expected)
	})

	t.Run("insufficient funds", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Bitcoin(10))

		err := wallet.Withdraw(Bitcoin(15))

		assertError(t, err, ErrInsufficientFunds)
		assertBalance(t, wallet, Bitcoin(10))
	})

}

func assertBalance(t testing.TB, wallet Wallet, expected Bitcoin) {
	t.Helper()

	balance := wallet.Balance()

	if balance != expected {
		t.Errorf("expected %s, got %s", expected, balance)
	}
}

func assertError(t testing.TB, err error, expected error) {
	t.Helper()

	if err == nil {
		t.Fatal("expected an error but got nil")
	}

	if !errors.Is(err, expected) {
		t.Errorf("expected error: %v, but got error: %v", expected, err)
	}
}

func assertNoError(t testing.TB, err error) {
	t.Helper()

	if err != nil {
		t.Fatal("expected nil but got an error")
	}
}
