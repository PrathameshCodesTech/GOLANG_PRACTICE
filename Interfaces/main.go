// An interface defines a set of method signatures.
// Any type that implements those methods is said to satisfy the interface.

package main

import (
	"fmt"
	"math"
)

// Define an interface
type Shape interface {
	Area() float64
	Perimeter() float64
}

// Struct 1: Circle
type Circle struct {
	Radius float64
}

// Struct 2: Rectangle
type Rectangle struct {
	Width, Height float64
}

// Implement Shape interface for Circle
func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.Radius
}

// Implement Shape interface for Rectangle
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

// Main function
func main() {
	var s Shape

	s = Circle{Radius: 5}
	fmt.Println("Circle Area:", s.Area())
	fmt.Println("Circle Perimeter:", s.Perimeter())

	s = Rectangle{Width: 4, Height: 3}
	fmt.Println("Rectangle Area:", s.Area())
	fmt.Println("Rectangle Perimeter:", s.Perimeter())
}


// Shape is an interface with two methods: Area() and Perimeter().

// Circle and Rectangle are structs that implement those methods.

// The variable s can hold any type that implements the Shape interface.