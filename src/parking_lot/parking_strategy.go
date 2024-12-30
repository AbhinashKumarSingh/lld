package parking_lot

type ParkingStrategy interface {
	FindSpot(spots []ParkingSpot, vehicleType string) ParkingSpot
}

// NearestSpaceStrategy for parking in the nearest available space
type NearestSpaceStrategy struct{}

func (s *NearestSpaceStrategy) FindSpot(spots []ParkingSpot) ParkingSpot {
	for _, spot := range spots {
		if spot.IsAvailable() {
			return spot
		}
	}
	return nil
}

// NearToElevatorStrategy for parking near the elevator (if applicable)
type NearToElevatorStrategy struct{}

func (s *NearToElevatorStrategy) FindSpot(spots []ParkingSpot) ParkingSpot {
	for _, spot := range spots {
		if spot.IsAvailable() {
			switch spot := spot.(type) {
			case *TwoWheelerParkingSpot:
				if spot.NearElevator {
					return spot
				}
			}
		}
	}
	return nil
}

// DefaultStrategy for finding the first available spot
// type DefaultStrategy struct{}

// func (s *DefaultStrategy) FindSpot(spots []ParkingSpot) ParkingSpot {
// 	for _, spot := range spots {
// 		if spot.IsAvailable() {
// 			return spot
// 		}
// 	}
// 	return nil
// }

type DefaultStrategy struct{}

func (d *DefaultStrategy) FindSpot(spots []ParkingSpot, vehicleType string) ParkingSpot {
	// Default logic to find the first available spot
	for _, spot := range spots {
		if spot.IsAvailable() {
			if vehicleType == "TwoWheeler" && spot.(*TwoWheelerParkingSpot) != nil {
				return spot
			}
			if vehicleType == "FourWheeler" && spot.(*FourWheelerParkingSpot) != nil {
				return spot
			}
		}
	}
	return nil
}

// NearestStrategy to find the nearest available spot (for example)
type NearestStrategy struct{}

func (n *NearestStrategy) FindSpot(spots []ParkingSpot, vehicleType string) ParkingSpot {
	// Logic to find the nearest available spot based on some criteria (example, first spot)
	// Can implement distance-based logic
	for _, spot := range spots {
		if spot.IsAvailable() {
			if vehicleType == "TwoWheeler" && spot.(*TwoWheelerParkingSpot) != nil {
				return spot
			}
			if vehicleType == "FourWheeler" && spot.(*FourWheelerParkingSpot) != nil {
				return spot
			}
		}
	}
	return nil
}
