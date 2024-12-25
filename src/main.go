package main

import (
	"fmt"
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
}
