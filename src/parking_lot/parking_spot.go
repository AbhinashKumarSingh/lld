package parking_lot

type ParkingSpot interface {
	GetID() int64
	IsAvailable() bool
	Park(vehicle *Vehicle) bool
	RemoveVehicle()
	GetPrice() int64
	GetVehicle() *Vehicle
}

// TwoWheelerParkingSpot struct
type TwoWheelerParkingSpot struct {
	ID           int64
	IsEmpty      bool
	Vehicle      *Vehicle
	Price        int64
	NearElevator bool
}

// Implement ParkingSpot methods for TwoWheelerParkingSpot
func (s *TwoWheelerParkingSpot) GetID() int64         { return s.ID }
func (s *TwoWheelerParkingSpot) IsAvailable() bool    { return s.IsEmpty }
func (s *TwoWheelerParkingSpot) GetPrice() int64      { return s.Price }
func (s *TwoWheelerParkingSpot) GetVehicle() *Vehicle { return s.Vehicle }
func (s *TwoWheelerParkingSpot) Park(vehicle *Vehicle) bool {
	if s.IsEmpty && vehicle.Type == "TwoWheeler" {
		s.Vehicle = vehicle
		s.IsEmpty = false
		return true
	}
	return false
}
func (s *TwoWheelerParkingSpot) RemoveVehicle() {
	s.Vehicle = nil
	s.IsEmpty = true
}

// FourWheelerParkingSpot struct
type FourWheelerParkingSpot struct {
	ID      int64
	IsEmpty bool
	Vehicle *Vehicle
	Price   int64
}

// Implement ParkingSpot methods for FourWheelerParkingSpot
func (s *FourWheelerParkingSpot) GetID() int64         { return s.ID }
func (s *FourWheelerParkingSpot) IsAvailable() bool    { return s.IsEmpty }
func (s *FourWheelerParkingSpot) GetPrice() int64      { return s.Price }
func (s *FourWheelerParkingSpot) GetVehicle() *Vehicle { return s.Vehicle }
func (s *FourWheelerParkingSpot) Park(vehicle *Vehicle) bool {
	if s.IsEmpty && vehicle.Type == "FourWheeler" {
		s.Vehicle = vehicle
		s.IsEmpty = false
		return true
	}
	return false
}
func (s *FourWheelerParkingSpot) RemoveVehicle() {
	s.Vehicle = nil
	s.IsEmpty = true
}
