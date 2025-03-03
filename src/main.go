package main

import (
	"fmt"
	"lld/src/coupon_discount"
	"lld/src/short_url"
)

func main() {
	// employee := srp.Employee{Name: "Alice", Salary: 50000}
	// fmt.Println(employee.GetDetails())

	// fileManager := srp.EmployeeFileManager{}
	// fileManager.SaveToFile(employee)

	// r := lsp.Rectangle{Weight: 10, Length: 20}

	// s := lsp.Square{Side: 12}

	// lsp.PrintArea(r)
	// lsp.PrintArea(s)

	// payment := &strategy_design_pattern.PaymentStrategy{}

	// payment.SetStrategy(&strategy_design_pattern.StripePayment{StripeAmount: 100})
	// payment.ProcessPayment(100)

	// payment.SetStrategy(&strategy_design_pattern.PaypalPayment{PaypalAmount: 101})
	// payment.ProcessPayment(101)

	// weather := &observer_design_pattern.Weather{}

	// phone := &observer_design_pattern.PhoneDisplay{ID: "1"}
	// desk := &observer_design_pattern.DesktopDisplay{ID: "2"}

	// weather.AddOberserver(phone)
	// weather.AddOberserver(desk)

	// weather.SetTemp(12.0)
	// weather.RemoveObserver(desk)
	// weather.SetTemp(24.0)

	// var coffee decorator_design_pattern.Coffee = &decorator_design_pattern.SimpleCoffee{}
	// fmt.Printf("%s : $%d\n", coffee.Description(), coffee.Cost())

	// coffee = decorator_design_pattern.NewMilkDecorator(coffee)
	// fmt.Printf("%s : $%d\n", coffee.Description(), coffee.Cost())
	// coffee = decorator_design_pattern.NewSugarDecorator(coffee)
	// fmt.Printf("%s : $%d\n", coffee.Description(), coffee.Cost())
	// coffee = decorator_design_pattern.NewWhippedCreamDecorator(coffee)

	// fmt.Printf("%s : $%d\n", coffee.Description(), coffee.Cost())

	// // Request a car
	// car, err := factory_design_pattern.VehicleFactory("car")
	// if err != nil {
	// 	fmt.Println("Error:", err)
	// 	return
	// }
	// fmt.Println(car.Drive())

	// // Request a bike
	// bike, err := factory_design_pattern.VehicleFactory("bike")
	// if err != nil {
	// 	fmt.Println("Error:", err)
	// 	return
	// }
	// fmt.Println(bike.Drive())

	// // Request an unknown vehicle type
	// _, err = factory_design_pattern.VehicleFactory("plane")
	// if err != nil {
	// 	fmt.Println("Error:", err)
	// 	// return
	// }
	// fmt.Printf("No ticket associated with this exit gate.")
	// // fmt.Println(unknown.Drive())

	// // Create parking spots for two-wheelers and four-wheelers
	// parkingTypeManager := &parking_lot.ParkingTypeManager{}

	// // Create a ParkingSpotManager instance
	// parkingSpotManager := &parking_lot.ParkingSpotManager{}

	// // Create and add parking spaces for two-wheelers and four-wheelers

	// twoWheelerSpot := &parking_lot.TwoWheelerParkingSpot{
	// 	ID:      1,
	// 	IsEmpty: true,
	// }
	// fourWheelerSpot := &parking_lot.FourWheelerParkingSpot{ID: 2, IsEmpty: true}

	// if twoWheelerSpot != nil {
	// 	parkingSpotManager.AddParkingSpace(twoWheelerSpot)
	// }
	// if fourWheelerSpot != nil {
	// 	parkingSpotManager.AddParkingSpace(fourWheelerSpot)
	// }

	// // Create an entrance gate
	// entranceGate := &parking_lot.EntranceGate{
	// 	GateNumber:         1,
	// 	ParkingSpotFactory: parkingTypeManager,
	// 	ParkingSpotManager: parkingSpotManager,
	// }

	// // Create a two-wheeler vehicle and book a parking spot
	// twoWheeler := &parking_lot.Vehicle{Type: "TwoWheeler", Plate: "TW1234"}
	// twoWheelerTicket, err := entranceGate.BookSpot(twoWheeler)
	// if err != nil {
	// 	fmt.Println("Error while booking spot:", err)
	// 	return
	// }

	// // Create exit gate to process exit and payment
	// exitGate := parking_lot.NewExitGate(1, nil, nil)
	// exitGate.SetTicket(twoWheelerTicket)

	// // Process the exit for the two-wheeler
	// fmt.Println("Processing exit for TwoWheeler...")
	// exitGate.ProcessExit()

	// // Create a four-wheeler vehicle and book a parking spot
	// fourWheeler := &parking_lot.Vehicle{Type: "FourWheeler", Plate: "FW5678"}
	// fourWheelerTicket, err := entranceGate.BookSpot(fourWheeler)
	// if err != nil {
	// 	fmt.Println("Error while booking spot:", err)
	// 	return
	// }

	// // Process the exit and payment for the four-wheeler
	// exitGate.SetTicket(fourWheelerTicket)
	// fmt.Println("Processing exit for FourWheeler...")
	// exitGate.ProcessExit()

	// game := tic_tac_toe.NewGame(3, []string{"Abhi", "Shiv"})

	// game.StartGame()

	p1 := &coupon_discount.Product{Name: "a", Price: 1000}
	p2 := &coupon_discount.Product{Name: "b", Price: 2000}

	shoppingCart := &coupon_discount.ShoppingCart{}
	shoppingCart.AddToCartPercent(*p1)
	shoppingCart.AddToCartValue(*p2)

	fmt.Printf("%f", shoppingCart.GetTotalPrice())

	short_url.Init()

}
