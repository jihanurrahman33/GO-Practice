package payment

func Checkout(payment PaymentProcessor, amount float32) {
	payment.Pay(amount)

}

type PaymentProcessor interface {
	Pay(amount float32)
}
