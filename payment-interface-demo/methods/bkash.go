package methods

import (
	"fmt"
	"interface_practice/payment"
)

type Bkash struct {
	gateWay payment.PaymentProcessor
}

func (b *Bkash) Pay(amount float32) {
	fmt.Println("Paid using Bkash", amount)
}
