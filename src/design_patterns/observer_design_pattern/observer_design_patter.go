package observer_design_pattern

import "fmt"

type Observer interface {
	Update(temp float64)
}

type Observable interface {
	AddOberserver(obj Observer)
	RemoveObserver(obj Observer)
	Notify()
}

type Weather struct {
	Observers []Observer
	Temp      float64
}

func (ob *Weather) AddOberserver(obj Observer) {
	ob.Observers = append(ob.Observers, obj)
}

func (ob *Weather) RemoveObserver(obj Observer) {
	for i, v := range ob.Observers {
		if v == obj {
			ob.Observers = append(ob.Observers[:i], ob.Observers[i+1:]...)
			break
		}
	}
}

func (ob *Weather) Notify() {
	for _, observer := range ob.Observers {
		observer.Update(ob.Temp)
	}
}

func (ob *Weather) SetTemp(t float64) {
	ob.Temp = t
	ob.Notify()
}

type PhoneDisplay struct {
	ID string
}

func (ob *PhoneDisplay) Update(t float64) {
	fmt.Printf("PhoneDisplay %s: Temperature updated to %.2f°C\n", ob.ID, t)
}

type DesktopDisplay struct {
	ID string
}

func (ob *DesktopDisplay) Update(t float64) {
	fmt.Printf("DesktopDisplay %s: Temperature updated to %.2f°C\n", ob.ID, t)
}
