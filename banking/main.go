package main

import (
	"banking/account"
	"fmt"
)

func main() {
	fmt.Println("--- Bank Account System ---")

	savAcc := account.SavingsAccount{
		Account: account.Account{
			AccountNumber: "SAV001",
			Balance:       1000.00,
			OwnerName:     "John Doe",
		},
		InterestRate: 0.02,
	}

	fmt.Println("\n--- Savings Account Operations ---")
	fmt.Println(savAcc.Account.String())

	err := savAcc.Deposit(200.00)
	if err != nil {
		fmt.Printf("Error depositing $%.2f to savings account. %+v\n", 200.00, err)
	}

	savAcc.AddInterest()
	err = savAcc.Withdraw(50.00)
	if err != nil {
		fmt.Println("Error:", err)
	}

	fmt.Println("Final Savings Details:", savAcc.Account.String())

	ovdAcc := account.OverdraftAccount{
		Account: account.Account{
			AccountNumber: "OVD002",
			Balance:       100.00,
			OwnerName:     "Nishak",
		},
		OverdraftLimit: 200.00,
	}

	fmt.Println("--- Overdraft account operations ---")
	fmt.Println(ovdAcc.Account.String())

	err = ovdAcc.Deposit(50.00)
	if err != nil {
		fmt.Printf("Error depositing $%.2f to savings account. %+v\n", 200.00, err)
	}

	err = ovdAcc.Withdraw(200.00)
	if err != nil {
		fmt.Println("Error:", err)
	}
	err = ovdAcc.Withdraw(100.00)
	if err != nil {
		fmt.Println("Error:", err)
	}
	err = ovdAcc.Withdraw(100.00)
	if err != nil {
		fmt.Println("Error:", err)
	}

	fmt.Println("Final Overdraft Details:", ovdAcc.Account.String())
}
