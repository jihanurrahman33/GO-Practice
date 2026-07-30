package account

import (
	"errors"
	"fmt"
)

type Account struct {
	AccountNumber string
	Balance       float64
	OwnerName     string
}

func (acc *Account) Deposit(amount float64) error {
	if amount <= 0 {
		return errors.New("Deposit amount must be positive")
	}
	acc.Balance += amount

	fmt.Printf("Deposited $%.2f to %s. New Balance: $%.2f\n", amount, acc.AccountNumber, acc.Balance)
	return nil
}

func (acc *Account) Withdraw(amount float64) error {
	if amount <= 0 {
		return errors.New("Withdrawal amount must be positive")
	}
	if acc.Balance < amount {
		return fmt.Errorf("Insufficient funds in %s. Balance: $%.2f, Tried to withdraw: $%.2f", acc.AccountNumber, acc.Balance, amount)
	}

	acc.Balance -= amount
	fmt.Printf("Withdrew $%.2f from %s. New Balance $%.2f\n", amount, acc.AccountNumber, acc.Balance)
	return nil
}

func (acc *Account) GetBalance() float64 {
	return acc.Balance
}

func (acc *Account) String() string {
	return fmt.Sprintf("Account [%s] Owner: %s, Balance: $%.2f", acc.AccountNumber, acc.OwnerName, acc.Balance)
}

type SavingsAccount struct {
	Account
	InterestRate float64
}

func (sa *SavingsAccount) AddInterest() {
	interest := sa.Balance * sa.InterestRate
	fmt.Printf("Adding interest $%.2f to savings account %s.", interest, sa.AccountNumber)
	err := sa.Deposit(interest)
	if err != nil {
		fmt.Printf("AddInterest: Error depositing $%.2f to savings account. %+v\n", interest, err)
	}
}

type OverdraftAccount struct {
	Account
	OverdraftLimit float64
}

func (oa *OverdraftAccount) Withdraw(amount float64) error {
	if amount <= 0 {
		return errors.New("withdrawal amount must be positive")
	}

	if (oa.Balance + oa.OverdraftLimit) < amount {
		return fmt.Errorf("withdrawal of $%.2f exceeds overdraft limit for %s. Available including overdraft: $%.2f", amount, oa.AccountNumber, oa.Balance+oa.OverdraftLimit)
	}
	oa.Balance -= amount
	fmt.Printf("Withdrew $%.2f from overaft account %s. New Balance: $%.2f\n", amount, oa.AccountNumber, oa.Balance)
	return nil
}
