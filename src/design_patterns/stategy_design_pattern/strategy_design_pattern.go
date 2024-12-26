package strategy_design_pattern

import "fmt"

type Payment interface {
	Pay(amount int)
}

type StripePayment struct {
	StripeAmount int
}

func (st *StripePayment) Pay(amount int) {
	fmt.Println("using stripe payment:", amount)
}

type PaypalPayment struct {
	PaypalAmount int
}

func (pp *PaypalPayment) Pay(amount int) {
	fmt.Println("using paypal payment:", amount)
}

type PaymentStrategy struct {
	Payment
}

func (ps *PaymentStrategy) SetStrategy(p Payment) {
	ps.Payment = p
}

func (ps *PaymentStrategy) ProcessPayment(amount int) {
	ps.Payment.Pay(amount)
}
