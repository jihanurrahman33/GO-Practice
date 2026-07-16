package main

import (
	account "bank_account/account"
	"fmt"
)

func main() {
	acc1 := account.BankAccount{
		Owner: "Nishak",
	}

	acc1.Deposit(100, 200, 300, 400, 500)
	fmt.Println("Balance:", acc1.Balance)
	acc1.Withdraw(500)
	fmt.Println("Balance:", acc1.Balance)
	acc1.ShowHistory()
}
