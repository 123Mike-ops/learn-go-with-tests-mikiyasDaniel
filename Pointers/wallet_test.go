package main

import "testing"

type Wallet struct{}

func (w Wallet) Deposit(amount int) {

}

func (w Wallet) Balance() int {
	return 0
}
func TestWallet(t *testing.T) {

	wallet := Wallet{}

	wallet.Deposit(10)

	got := wallet.Balance()
	want := 10

	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}
