package parking_lot

import "fmt"

// PaymentStrategy defines the common interface for different payment methods
type PaymentStrategy interface {
	Pay(ticket *Ticket) bool // Process payment for the ticket
}

// PayPalPayment implements the PaymentStrategy interface for PayPal payments
type PayPalPayment struct {
	ClientID string
	Secret   string
}

// NewPayPalPayment creates a new instance of PayPalPayment with the credentials
func NewPayPalPayment(clientID, secret string) *PayPalPayment {
	return &PayPalPayment{ClientID: clientID, Secret: secret}
}

// Pay processes the payment for the ticket using PayPal
func (p *PayPalPayment) Pay(ticket *Ticket) bool {
	// Simulate PayPal payment process
	fmt.Printf("Processing payment of $%d for ticket %d via PayPal.\n", ticket.Price, ticket.TicketID)

	return true
}

type StripePayment struct {
	APIKey string
}

// NewStripePayment creates a new instance of StripePayment with the API Key
func NewStripePayment(apiKey string) *StripePayment {
	return &StripePayment{APIKey: apiKey}
}

// Pay processes the payment for the ticket using Stripe
func (s *StripePayment) Pay(ticket *Ticket) bool {

	fmt.Printf("Processing payment of $%d for ticket %d via Stripe.\n", ticket.Price, ticket.TicketID)

	return true
}
