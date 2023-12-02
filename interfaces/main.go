package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("This is the example for better way to write interfaces...")

	c := circle{radius: 10}
	fmt.Printf("Area of a circle is %s and perimeter is %s", c.area(), c.perimeter())
	s := square{side: 3.4}
	fmt.Printf("Area of a circle is %s and perimeter is %s", s.area(), s.perimeter())
}

type circle struct {
	radius float64
}
type square struct {
	side float64
}

func (c circle) area() float64 {
	return 3.14 * c.radius * c.radius
}

func (c circle) perimeter() float64 {
	return 2 * 3.14 * c.radius
}
func (s square) area() float64 {
	return math.Pow(float64(s.side), 2)
}

func (s square) perimeter() float64 {
	return 4 * s.side
}
