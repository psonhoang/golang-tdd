package main

import (
	"errors"
	"fmt"
)

type Bitcoin int

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

type Wallet struct {
	balance	Bitcoin
}

var ErrInsuficientFunds = errors.New("cannot withdraw, insufficient funds")

func (w *Wallet) Deposit(amount Bitcoin) {
	// fmt.Printf("the address of balance in Deposit() is %p \n", &w.balance)
	w.balance += amount
}

func (w *Wallet) Withdraw(amount Bitcoin) error {
	if amount > w.balance {
		return ErrInsuficientFunds
	}

	w.balance -= amount
	return nil
}

func (w *Wallet) Balance() Bitcoin {
	
	return w.balance
}