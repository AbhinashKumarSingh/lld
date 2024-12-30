package parking_lot

type PricingStrategy interface {
	CalculateCost(ticket *Ticket) int64
}

// TwoWheelerPricingStrategy implements the PricingStrategy for two-wheelers
type TwoWheelerPricingStrategy struct{}

func (t *TwoWheelerPricingStrategy) CalculateCost(ticket *Ticket) int64 {
	// For example, assume 50 units per hour for two-wheeler parking
	// Here, we can customize it based on your logic (e.g., time spent, parking type, etc.)
	return ticket.ParkingSpot.GetPrice() * 50
}

// FourWheelerPricingStrategy implements the PricingStrategy for four-wheelers
type FourWheelerPricingStrategy struct{}

func (f *FourWheelerPricingStrategy) CalculateCost(ticket *Ticket) int64 {
	// For example, assume 100 units per hour for four-wheeler parking
	return ticket.ParkingSpot.GetPrice() * 100
}
