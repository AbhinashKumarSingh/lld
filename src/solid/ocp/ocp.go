package solid

import "fmt"

type Shape interface {
	Area() int64
	Perimeter() int64
}

type Circle struct {
	Radius int64
}

type Rectangle struct {
	Length  int64
	Breadth int64
}

func (c Circle) Area() int64 {
	return c.Radius * c.Radius
}

func (c Circle) Perimeter() int64 {
	return 2 * 3 * c.Radius
}
func (r Rectangle) Area() int64 {
	return r.Length * r.Breadth
}

func (r Rectangle) Perimeter() int64 {
	return 2 * (r.Length + r.Breadth)
}

func Print(shape Shape) {
	fmt.Printf("The area is: %d", shape.Area())
	fmt.Printf("The perimeter is: %d", shape.Perimeter())
}

// type Shape interface {
// 	Area() float64
// }

// // Circle struct implements Shape
// type Circle struct {
// 	Radius float64
// }

// func (c Circle) Area() float64 {
// 	return 3.14 * c.Radius * c.Radius
// }

// // Square struct implements Shape
// type Square struct {
// 	Side float64
// }

// func (s Square) Area() float64 {
// 	return s.Side * s.Side
// }

// // PrintArea prints the area of any shape
// func PrintArea(shape Shape) {
// 	fmt.Printf("The area is: %.2f\n", shape.Area())
// }
