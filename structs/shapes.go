package main

import "math"

// Shape interface for all objects
type Shape interface {
	Area() float64
}

// Rectangle contains width and height
type Rectangle struct {
	Width  float64
	Height float64
}

// Area for Rectangle
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

// Circle struct
type Circle struct {
	Radius float64
}

// Area returns area of Circle
func (c Circle) Area() float64 {
	return c.Radius * c.Radius * math.Pi
}

// Triangle struct
type Triangle struct {
	Base   float64
	Height float64
}

// Area returns area of Triangle
func (t Triangle) Area() float64 {
	return t.Base * t.Height * 0.5
}

// Perimeter returns perimeter of Rectangle
func Perimeter(rec Rectangle) float64 {
	return 2 * (rec.Width + rec.Height)
}
