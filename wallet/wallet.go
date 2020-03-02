package wallet

import (
	"errors"
	"fmt"
)

// ErrInsufficientFunds -
var ErrInsufficientFunds = errors.New("cannot withdraw, insufficient funds")

// Stringer --
type Stringer interface {
	String() string
}

// Bitcoin -
type Bitcoin int

// String --
func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

// Wallet -
type Wallet struct {
	balance Bitcoin
}

// Deposit --
func (w *Wallet) Deposit(amount Bitcoin) {
	w.balance += amount
}

// Withdraw -
func (w *Wallet) Withdraw(amount Bitcoin) error {
	if amount > w.balance {
		return ErrInsufficientFunds
	}

	w.balance -= amount
	return nil
}

// Balance --
func (w *Wallet) Balance() Bitcoin {
	return w.balance
}
