package parking_lot

import "fmt"

// CostComputeFactory is responsible for creating CostCompute objects based on vehicle type
type CostComputeFactory struct{}

// NewCostComputeFactory creates a new instance of CostComputeFactory
func NewCostComputeFactory() *CostComputeFactory {
	return &CostComputeFactory{}
}

// GetCostCompute returns the appropriate CostCompute based on the vehicle type in the ticket
func (factory *CostComputeFactory) GetCostCompute(ticket *Ticket) *CostCompute {
	switch ticket.Vehicle.Type {
	case "TwoWheeler":
		return &CostCompute{
			TwoWheelerStrategy: &TwoWheelerPricingStrategy{},
		}
	case "FourWheeler":
		return &CostCompute{
			FourWheelerStrategy: &FourWheelerPricingStrategy{},
		}
	default:
		fmt.Println("Unsupported vehicle type:", ticket.Vehicle.Type)
		return nil
	}
}
