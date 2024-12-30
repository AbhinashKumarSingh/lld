package parking_lot

import "fmt"

// ExitGate represents the exit gate of the parking lot
type ExitGate struct {
	ID                 int64
	Ticket             *Ticket
	CostComputeFactory *CostComputeFactory
	PaymentFactory     *PaymentStrategyFactory
}

// NewExitGate creates a new exit gate with a ticket reference
func NewExitGate(id int64, costFactory *CostComputeFactory, paymentFactory *PaymentStrategyFactory) *ExitGate {
	return &ExitGate{
		ID:                 id,
		CostComputeFactory: costFactory,
		PaymentFactory:     paymentFactory,
	}
}

// SetTicket associates a ticket with the exit gate
func (g *ExitGate) SetTicket(ticket *Ticket) {
	g.Ticket = ticket
}

// ProcessExit calculates the price and completes the exit process
func (g *ExitGate) ProcessExit() {
	if g.Ticket == nil {
		fmt.Println("No ticket associated with this exit gate.")
		return
	}

	fmt.Println("No ticket associated with this exit gate.")
	// Get the cost computation strategy based on the vehicle type in the ticket
	costCompute := g.CostComputeFactory.GetCostCompute(g.Ticket)

	// Calculate the final price using the pricing strategy
	price := g.CalculatePrice(g.Ticket, costCompute)

	// Update the ticket with the final price
	g.Ticket.Price = price

	// Generate the ticket details
	fmt.Println("Final Ticket Details at Exit:")
	fmt.Println(g.Ticket.GenerateTicket())

	paymentStrategy := g.PaymentFactory.GetPaymentStrategy(g.Ticket)

	// Process the payment using the selected strategy
	if paymentStrategy.Pay(g.Ticket) {
		fmt.Println("Payment Successful!")
	} else {
		fmt.Println("Payment Failed!")
	}
}

// CalculatePrice calculates the price using the provided CostCompute object
func (g *ExitGate) CalculatePrice(ticket *Ticket, costCompute *CostCompute) int64 {
	// Delegate the price calculation to CostCompute
	return costCompute.CalculateTicketCost(ticket)
}
