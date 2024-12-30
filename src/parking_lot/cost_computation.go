package parking_lot

import "fmt"

// CostCompute will manage which pricing strategy to use for cost calculation
type CostCompute struct {
	TwoWheelerStrategy  PricingStrategy
	FourWheelerStrategy PricingStrategy
}

// NewCostCompute creates an instance of CostCompute with pricing strategies for two-wheelers and four-wheelers
func NewCostCompute() *CostCompute {
	return &CostCompute{
		TwoWheelerStrategy:  &TwoWheelerPricingStrategy{},
		FourWheelerStrategy: &FourWheelerPricingStrategy{},
	}
}

// SetPricingStrategy sets the pricing strategy based on vehicle type
func (c *CostCompute) SetPricingStrategy(vehicleType string) PricingStrategy {
	switch vehicleType {
	case "TwoWheeler":
		return c.TwoWheelerStrategy
	case "FourWheeler":
		return c.FourWheelerStrategy
	default:
		fmt.Println("Unsupported vehicle type:", vehicleType)
		return nil
	}
}

// CalculateTicketCost calculates the cost of the parking ticket using the appropriate pricing strategy
func (c *CostCompute) CalculateTicketCost(ticket *Ticket) int64 {
	strategy := c.SetPricingStrategy(ticket.Vehicle.Type)
	if strategy == nil {
		fmt.Println("No valid pricing strategy found")
		return 0
	}
	return strategy.CalculateCost(ticket)
}
