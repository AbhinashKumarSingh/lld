package decorator_design_pattern

// Component: The Coffee interface
type Coffee interface {
	Cost() int
	Description() string
}

// Concrete Component: SimpleCoffee
type SimpleCoffee struct{}

func (s *SimpleCoffee) Cost() int {
	return 5 // Base price for coffee
}

func (s *SimpleCoffee) Description() string {
	return "Simple Coffee"
}

// Decorator: Base decorator struct
type CoffeeDecorator struct {
	coffee Coffee // Wrapping the Coffee interface
}

func (d *CoffeeDecorator) Cost() int {
	return d.coffee.Cost()
}

func (d *CoffeeDecorator) Description() string {
	return d.coffee.Description()
}

// Concrete Decorators
type MilkDecorator struct {
	*CoffeeDecorator
}

func NewMilkDecorator(c Coffee) Coffee {
	return &MilkDecorator{
		CoffeeDecorator: &CoffeeDecorator{coffee: c},
	}
}

func (m *MilkDecorator) Cost() int {
	return m.coffee.Cost() + 2 // Adding milk costs $2
}

func (m *MilkDecorator) Description() string {
	return m.coffee.Description() + ", Milk"
}

type SugarDecorator struct {
	*CoffeeDecorator
}

func NewSugarDecorator(c Coffee) Coffee {
	return &SugarDecorator{
		CoffeeDecorator: &CoffeeDecorator{coffee: c},
	}
}

func (s *SugarDecorator) Cost() int {
	return s.coffee.Cost() + 1 // Adding sugar costs $1
}

func (s *SugarDecorator) Description() string {
	return s.coffee.Description() + ", Sugar"
}

type WhippedCreamDecorator struct {
	*CoffeeDecorator
}

func NewWhippedCreamDecorator(c Coffee) Coffee {
	return &WhippedCreamDecorator{
		CoffeeDecorator: &CoffeeDecorator{coffee: c},
	}
}

func (w *WhippedCreamDecorator) Cost() int {
	return w.coffee.Cost() + 3 // Adding whipped cream costs $3
}

func (w *WhippedCreamDecorator) Description() string {
	return w.coffee.Description() + ", Whipped Cream"
}
