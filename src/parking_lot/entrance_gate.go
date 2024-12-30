package parking_lot

import (
	"fmt"
	"time"
)

type EntranceGate struct {
	GateNumber         int
	ParkingSpotFactory *ParkingTypeManager
	ParkingSpotManager *ParkingSpotManager
}

func (g *EntranceGate) FindParkingSpace(vehicleType string) ParkingSpot {
	return g.ParkingSpotManager.FindParkingSpace(vehicleType)
}

func (g *EntranceGate) BookSpot(vehicle *Vehicle) (*Ticket, error) {
	spot := g.FindParkingSpace(vehicle.Type)
	if spot == nil {
		return nil, fmt.Errorf("no available spot for vehicle type %s", vehicle.Type)
	}

	// Book the spot
	spot.Park(vehicle)

	// Generate ticket
	ticket := &Ticket{
		TicketID:    time.Now().Unix(),
		Vehicle:     *vehicle,
		ParkingSpot: spot,
		EntryTime:   time.Now().Format("2006-01-02"),
		// Price:       spot.GetPrice(),
	}
	return ticket, nil
}
