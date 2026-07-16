package account

import "fmt"

type BankAccount struct {
	Owner        string
	Balance      float64
	Transactions []float64
}

func (bAcc *BankAccount) Deposit(amounts ...float64) {
	if len(amounts) == 0 {
		fmt.Println("Please Input Valid Amounts")
		return
	}

	for _, amount := range amounts {
		bAcc.Balance += amount
		bAcc.Transactions = append(bAcc.Transactions, amount)
	}
	fmt.Println("Deposit Success")

}
func (bAcc *BankAccount) Withdraw(amount float64) {
	if amount <= 0 {
		fmt.Println("Withdrawal Amount must be greater than zero")
		return
	}

	if bAcc.Balance < amount {
		fmt.Println("Insufficient Balance")
		return
	}

	bAcc.Balance -= amount
	bAcc.Transactions = append(bAcc.Transactions, -amount)
	fmt.Printf("Withdraw Amount %f Successfull", amount)
	fmt.Println()
}
func (bAcc *BankAccount) ShowHistory() {
	fmt.Println("Transactions:")

	for _, t := range bAcc.Transactions {
		fmt.Println(t)
	}

	fmt.Println("Current Balance:", bAcc.Balance)
}
