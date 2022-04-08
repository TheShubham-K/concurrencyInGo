package main

import "errors"

type BitcoinAccount struct {
	bankAcc IBankAccount
	balance int
	fee     int
}

func NewBitcoinAccount() *BitcoinAccount {
	return &BitcoinAccount{
		balance: 0,
		fee:     300,
	}
}

func (b *BitcoinAccount) GetBalance() int {
	return b.balance
}

func (b *BitcoinAccount) Deposit(amount int) {
	b.balance += amount
}

func (b *BitcoinAccount) Withdraw(amount int) error {
	newBalance := b.balance - amount - b.fee
	if newBalance < 0 {
		return errors.New("Insufficient Balance")
	}
	b.balance = newBalance
	return nil
}
