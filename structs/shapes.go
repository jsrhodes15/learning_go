package main

import "math"

/**
 * Declaring interface to define functions that can be used by different types
 */

// Shape is implemented by anything that can tell us its Area
type Shape interface {
	Area() float64
}

/**
 * Declaring struct to create data types (bundle related data together)
 */

// Rectangle has the dimensions of a rectangle
type Rectangle struct {
	Width  float64
	Height float64
}

/**
 * Adding methods to add functionality to data types (implement an interface)
 * func (receiverName ReceiverType) MethodName(args)
 */

// Area returns the area of the rectangle
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

// Perimeter returns the perimeter of a rectangle
func Perimeter(rectangle Rectangle) float64 {
	return 2 * (rectangle.Width + rectangle.Height)
}

// Circle represents a circle...
type Circle struct {
	Radius float64
}

// Area returns the area of the circle
func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

// Triangle represents the dimensions of a triangle
type Triangle struct {
	Base   float64
	Height float64
}

// Area returns the area of the triangle
func (t Triangle) Area() float64 {
	return (t.Base * t.Height) * 0.5
}
