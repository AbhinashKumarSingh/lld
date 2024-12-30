package parking_lot

import "fmt"

type Ticket struct {
	TicketID  int64  `json:"id"`
	EntryTime string `json:"entry_time"`
	Vehicle
	ParkingSpot
	Price int64
}

func (t *Ticket) GenerateTicket() string {
	return fmt.Sprintf("TicketID: %d, Vehicle: %s, SpotID: %d, IssueTime: %s, ",
		t.TicketID, t.Vehicle.Plate, t.ParkingSpot.GetID(), string(t.EntryTime))
}
