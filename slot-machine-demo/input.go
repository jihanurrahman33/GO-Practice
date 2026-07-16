package main

import "fmt"

func GetName() string {
	name := ""

	fmt.Println("Welcome to Casino X...")
	fmt.Printf("Enter your Name: ")
	_, err := fmt.Scanln(&name)

	if err != nil {
		return ""
	}

	fmt.Printf("Welcome %s, let's play!\n", name)

	return name
}

func GetBet(balance uint) uint {
	var bet uint
	for true {
		fmt.Printf("Enter your bet,or 0 to quit (balance=$%d) : ", balance)
		fmt.Scan(&bet)

		if bet > balance {
			fmt.Println("Bet cannot be larger than your balance")
		} else {
			break
		}

	}
	return bet
}
