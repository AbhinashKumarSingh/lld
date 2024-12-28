package factory_design_pattern

import "fmt"

// Vehicle interface: Represents the product
type Vehicle interface {
	Drive() string
}

// Car: A concrete implementation of Vehicle
type Car struct{}

func (c *Car) Drive() string {
	return "Driving a Car"
}

// Bike: Another concrete implementation of Vehicle
type Bike struct{}

func (b *Bike) Drive() string {
	return "Riding a Bike"
}

// Factory function: Returns a Vehicle based on input
func VehicleFactory(vehicleType string) (Vehicle, error) {
	switch vehicleType {
	case "car":
		return &Car{}, nil
	case "bike":
		return &Bike{}, nil
	default:
		return nil, fmt.Errorf("unknown vehicle type: %s", vehicleType)
	}
}
