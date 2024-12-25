package isp

import "fmt"

// Segregated interfaces
type Workable interface {
	Work()
}

type Restable interface {
	Rest()
}

// Human implements both Workable and Restable
type Human struct{}

func (h Human) Work() {
	fmt.Println("Human is working")
}

func (h Human) Rest() {
	fmt.Println("Human is resting")
}

// Robot only implements Workable
type Robot struct{}

func (r Robot) Work() {
	fmt.Println("Robot is working")
}
