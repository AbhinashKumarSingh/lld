package parking_lot

import "fmt"

type ParkingTypeManager struct {
	TwoWheelerManager  *TwoWheelerParkingManager
	FourWheelerManager *FourWheelerParkingManager
}

// GetParkingManager returns the appropriate parking manager based on vehicle type
func (ptm *ParkingTypeManager) GetParkingManager(vehicleType string) interface{} {
	switch vehicleType {
	case "TwoWheeler":
		if ptm.TwoWheelerManager == nil {
			ptm.TwoWheelerManager = &TwoWheelerParkingManager{}
		}
		return ptm.TwoWheelerManager
	case "FourWheeler":
		if ptm.FourWheelerManager == nil {
			ptm.FourWheelerManager = &FourWheelerParkingManager{}
		}
		return ptm.FourWheelerManager
	default:
		fmt.Println("Unsupported vehicle type:", vehicleType)
		return nil
	}
}
