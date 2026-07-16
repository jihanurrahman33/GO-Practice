package main

type BankAccount struct {
	Balance float64
}

func (b *BankAccount) Deposit(amount *float64) {
	b.Balance += *amount
}
