package main

import (
	"fmt"
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

}
