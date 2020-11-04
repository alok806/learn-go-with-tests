package main

import "fmt"

/*
Stringer for Bitcoin
*/
type Stringer interface {
	String() string
}

/*
Bitcoin type int
*/
type Bitcoin int

/*
Wallet struct
*/
type Wallet struct {
	balance Bitcoin
}

/*
String print for Bitcoin
*/
func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

/*
Deposit amount to wallet
*/
func (w *Wallet) Deposit(amount Bitcoin) {
	w.balance += amount
}

/*
Balance check for wallet
*/
func (w *Wallet) Balance() Bitcoin {
	return w.balance
}

/*
Withdraw from wallet
*/
func (w *Wallet) Withdraw(amount Bitcoin) {
	w.balance -= amount
}
