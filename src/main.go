package main

import (
	"fmt"
	"lld/src/design_patterns/decorator_design_pattern"
	"lld/src/design_patterns/factory_design_pattern"
	"lld/src/design_patterns/observer_design_pattern"
	strategy_design_pattern "lld/src/design_patterns/stategy_design_pattern"
	"lld/src/solid/lsp"
	"lld/src/solid/srp" // Import the srp package
)

func main() {
	employee := srp.Employee{Name: "Alice", Salary: 50000}
	fmt.Println(employee.GetDetails())

	fileManager := srp.EmployeeFileManager{}
	fileManager.SaveToFile(employee)

	r := lsp.Rectangle{Weight: 10, Length: 20}

	s := lsp.Square{Side: 12}

	lsp.PrintArea(r)
	lsp.PrintArea(s)

	payment := &strategy_design_pattern.PaymentStrategy{}

	payment.SetStrategy(&strategy_design_pattern.StripePayment{StripeAmount: 100})
	payment.ProcessPayment(100)

	payment.SetStrategy(&strategy_design_pattern.PaypalPayment{PaypalAmount: 101})
	payment.ProcessPayment(101)

	weather := &observer_design_pattern.Weather{}

	phone := &observer_design_pattern.PhoneDisplay{ID: "1"}
	desk := &observer_design_pattern.DesktopDisplay{ID: "2"}

	weather.AddOberserver(phone)
	weather.AddOberserver(desk)

	weather.SetTemp(12.0)
	weather.RemoveObserver(desk)
	weather.SetTemp(24.0)

	var coffee decorator_design_pattern.Coffee = &decorator_design_pattern.SimpleCoffee{}
	fmt.Printf("%s : $%d\n", coffee.Description(), coffee.Cost())

	coffee = decorator_design_pattern.NewMilkDecorator(coffee)
	fmt.Printf("%s : $%d\n", coffee.Description(), coffee.Cost())
	coffee = decorator_design_pattern.NewSugarDecorator(coffee)
	fmt.Printf("%s : $%d\n", coffee.Description(), coffee.Cost())
	coffee = decorator_design_pattern.NewWhippedCreamDecorator(coffee)

	fmt.Printf("%s : $%d\n", coffee.Description(), coffee.Cost())

	// Request a car
	car, err := factory_design_pattern.VehicleFactory("car")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println(car.Drive())

	// Request a bike
	bike, err := factory_design_pattern.VehicleFactory("bike")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println(bike.Drive())

	// Request an unknown vehicle type
	unknown, err := factory_design_pattern.VehicleFactory("plane")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println(unknown.Drive())
}
