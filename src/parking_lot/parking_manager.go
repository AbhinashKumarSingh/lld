package parking_lot

import "fmt"

// type ParkingSpotManager struct {
// 	Spots []ParkingSpot
// }

// // AddParkingSpace adds a new parking spot
// func (m *ParkingSpotManager) AddParkingSpace(spot ParkingSpot) {
// 	m.Spots = append(m.Spots, spot)
// }

// // RemoveParkingSpace removes a parking spot by ID
// func (m *ParkingSpotManager) RemoveParkingSpace(id int64) {
// 	for i, spot := range m.Spots {
// 		if spot.GetID() == id {
// 			m.Spots = append(m.Spots[:i], m.Spots[i+1:]...)
// 			break
// 		}
// 	}
// }

// // FindParkingSpace finds the first available parking spot
// func (m *ParkingSpotManager) FindParkingSpace(vehicleType string) ParkingSpot {
// 	for _, spot := range m.Spots {
// 		if spot.IsAvailable() {
// 			switch vehicleType {
// 			case "TwoWheeler":
// 				if _, ok := spot.(*TwoWheelerParkingSpot); ok {
// 					return spot
// 				}
// 			case "FourWheeler":
// 				if _, ok := spot.(*FourWheelerParkingSpot); ok {
// 					return spot
// 				}
// 			}
// 		}
// 	}
// 	return nil
// }

// // Specialized TwoWheelerParkingManager
// type TwoWheelerParkingManager struct {
// 	ParkingSpotManager
// }

// func (m *TwoWheelerParkingManager) AddVehicle(vehicle *Vehicle) bool {
// 	if vehicle.Type != "TwoWheeler" {
// 		fmt.Println("Invalid vehicle type for TwoWheelerParkingManager")
// 		return false
// 	}
// 	spot := m.FindParkingSpace(vehicle.Type)
// 	if spot == nil {
// 		fmt.Println("No available spot for TwoWheeler")
// 		return false
// 	}
// 	return spot.Park(vehicle)
// }

// // Specialized FourWheelerParkingManager
// type FourWheelerParkingManager struct {
// 	ParkingSpotManager
// }

// func (m *FourWheelerParkingManager) AddVehicle(vehicle *Vehicle) bool {
// 	if vehicle.Type != "FourWheeler" {
// 		fmt.Println("Invalid vehicle type for FourWheelerParkingManager")
// 		return false
// 	}
// 	spot := m.FindParkingSpace(vehicle.Type)
// 	if spot == nil {
// 		fmt.Println("No available spot for FourWheeler")
// 		return false
// 	}
// 	return spot.Park(vehicle)
// }

type ParkingSpotManager struct {
	Spots    []ParkingSpot
	Strategy ParkingStrategy
}

// AddParkingSpace adds a new parking spot to the manager
func (m *ParkingSpotManager) AddParkingSpace(spot ParkingSpot) {
	m.Spots = append(m.Spots, spot)
}

// FindParkingSpace uses the current strategy to find a parking space
func (m *ParkingSpotManager) FindParkingSpace(vehicleType string) ParkingSpot {
	if m.Strategy == nil {
		m.Strategy = &DefaultStrategy{} // Default strategy if none is set
	}
	return m.Strategy.FindSpot(m.Spots, vehicleType)
}

// SetStrategy allows dynamic setting of the parking strategy
func (m *ParkingSpotManager) SetStrategy(strategy ParkingStrategy) {
	m.Strategy = strategy
}

// Specialized TwoWheelerParkingManager for managing two-wheeler spots
type TwoWheelerParkingManager struct {
	ParkingSpotManager
}

// AddVehicle adds a two-wheeler vehicle to the parking spot
func (m *TwoWheelerParkingManager) AddVehicle(vehicle *Vehicle) bool {
	if vehicle.Type != "TwoWheeler" {
		fmt.Println("Invalid vehicle type for TwoWheelerParkingManager")
		return false
	}
	spot := m.FindParkingSpace(vehicle.Type)
	if spot == nil {
		fmt.Println("No available spot for TwoWheeler")
		return false
	}
	return spot.Park(vehicle)
}

// Specialized FourWheelerParkingManager for managing four-wheeler spots
type FourWheelerParkingManager struct {
	ParkingSpotManager
}

// AddVehicle adds a four-wheeler vehicle to the parking spot
func (m *FourWheelerParkingManager) AddVehicle(vehicle *Vehicle) bool {
	if vehicle.Type != "FourWheeler" {
		fmt.Println("Invalid vehicle type for FourWheelerParkingManager")
		return false
	}
	spot := m.FindParkingSpace(vehicle.Type)
	if spot == nil {
		fmt.Println("No available spot for FourWheeler")
		return false
	}
	return spot.Park(vehicle)
}
