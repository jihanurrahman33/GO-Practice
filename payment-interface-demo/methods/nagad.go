package methods

import (
	"fmt"
	"interface_practice/payment"
)

type Nagad struct {
	gateWay payment.PaymentProcessor
}

func (n *Nagad) Pay(amount float32) {
	fmt.Println("Paid Using Nagad", amount)
}
