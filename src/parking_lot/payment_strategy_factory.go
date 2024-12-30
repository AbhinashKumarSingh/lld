package parking_lot

// PaymentStrategyFactory is used to return the correct payment strategy
type PaymentStrategyFactory struct{}

// NewPaymentStrategyFactory creates a new PaymentStrategyFactory
func NewPaymentStrategyFactory() *PaymentStrategyFactory {
	return &PaymentStrategyFactory{}
}

// GetPaymentStrategy returns the appropriate payment strategy based on the vehicle type or any other context
func (psf *PaymentStrategyFactory) GetPaymentStrategy(ticket *Ticket) PaymentStrategy {
	// For simplicity, return PayPal for four-wheelers and Stripe for two-wheelers
	// This can be extended to use user preferences, payment gateways, etc.
	if ticket.Vehicle.Type == "FourWheeler" {
		return NewPayPalPayment("payPalClientID", "payPalSecret")
	}
	return NewStripePayment("stripeAPIKey")
}
