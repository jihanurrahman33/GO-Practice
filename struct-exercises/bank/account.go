package bank

import "fmt"

type bankAccount struct {
	accountHolder string
	accountNumber string
	balance       float64
}

func (bAcc *bankAccount) Deposit(amount float64) {
	if amount <= 0 {
		fmt.Println("Deposit Amount must have to be greater than zero")
		return
	}
	bAcc.balance += amount

}
func (bAcc *bankAccount) Withdraw(amount float64) {
	if amount <= 0 {
		fmt.Println("Withdraw Amount must have to be greater than zero")
		return
	}
	if bAcc.balance < amount {
		fmt.Println("You dont have enough balance in your acc")
		return
	}
	bAcc.balance -= amount
}
func (bAcc *bankAccount) CheckBalance() {
	fmt.Println("Available Balance:", bAcc.balance)
}
