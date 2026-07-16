package methods

import (
	"fmt"
	"interface_practice/payment"
)

type Rocket struct {
	gateWay payment.PaymentProcessor
}

func (n *Rocket) Pay(amount float32) {
	fmt.Println("Paid Using Rocket", amount)
}
