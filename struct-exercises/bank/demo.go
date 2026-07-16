package bank

func RegisterBank() {
	user1 := bankAccount{
		accountHolder: "Md Jihanur Rahman",
		accountNumber: "10023576843957",
		balance:       0.00,
	}
	user2 := bankAccount{
		accountHolder: "Jr Nishak",
		accountNumber: "10034834738685",
		balance:       0.00,
	}

	user1.Deposit(100)
	user1.Withdraw(20)
	user1.CheckBalance()
	user2.Deposit(200)
	user2.Withdraw(70)
	user2.CheckBalance()
}
