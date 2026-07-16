package main

import (
	methods "interface_practice/methods"
	"interface_practice/payment"
)

func main() {
	// bkash := Bkash{}
	nagad := methods.Nagad{}
	// rocket := Rocket{}
	payment.Checkout(&nagad, 100)
}
