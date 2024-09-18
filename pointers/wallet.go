package main

import (
	"errors"
	"fmt"
)

// Bitcoin itself
type Bitcoin int

// Wallet for Bitcoin
type Wallet struct {
	balance Bitcoin
}

// Deposit money in wallet
func (w *Wallet) Deposit(amount Bitcoin) {
	w.balance += amount
}

// Balance of wallet
func (w *Wallet) Balance() Bitcoin {
	return w.balance
}

var ErrInsufficientFunds = errors.New("cannot withdraw, insufficient funds")

// Withdraw money from balance
func (w *Wallet) Withdraw(amount Bitcoin) error {

	if amount > w.balance {
		return ErrInsufficientFunds
	}

	w.balance -= amount
	return nil
}

type Stringer interface {
	String() string
}

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}
