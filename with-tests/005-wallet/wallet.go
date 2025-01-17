package wallet

import (
	"errors"
	"fmt"
)

type Bitcoin int

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

type Wallet struct {
	balance Bitcoin
}

func (w *Wallet) Withdraw(val Bitcoin) error {
	if w.balance < val {
		return errors.New("there is no suficient balance to withdraw")
	}
	w.balance -= val
	return nil
}

func (w *Wallet) Deposit(val Bitcoin) {
	w.balance += val
}

func (w *Wallet) Balance() Bitcoin {
	return w.balance
}
