package main

import "fmt"

func change(x *int) {
	*x = 100
}

func main() {

	num := 10
	change(&num)
	fmt.Println(num)

	user := user{"Nishak", 23}
	birthday(&user)
	fmt.Println(user.Age)

	acc1 := BankAccount{
		Balance: 100.0,
	}
	x := 100.00
	acc1.Deposit(&x)

	fmt.Println(acc1.Balance)
	a, b := 10, 20
	newA := swap(a, &b)

	fmt.Println(newA, b)
}
