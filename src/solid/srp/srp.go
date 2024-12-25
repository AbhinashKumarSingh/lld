package srp

import "fmt"

type Employee struct {
	Name   string
	Salary float64
}

// Responsibility: Employee details
func (e Employee) GetDetails() string {
	return fmt.Sprintf("Name: %s, Salary: %.2f", e.Name, e.Salary)
}

type EmployeeFileManager struct{}

// Responsibility: Save employee data to a file
func (efm EmployeeFileManager) SaveToFile(employee Employee) {
	fmt.Println("Saving employee to file...")
	// File writing logic
}

func main() {
	employee := Employee{Name: "Alice", Salary: 50000}
	fmt.Println(employee.GetDetails())

	fileManager := EmployeeFileManager{}
	fileManager.SaveToFile(employee)
}
