package shapes

import "math"

// Interfaces (with methods)
type Shape interface {
	Area() float64
}

// Structs (with embedded types)
type Rectangle struct {
	Width  float64
	Height float64
}

type Circle struct {
	Radius float64
}

type Triangle struct {
	Base   float64
	Height float64
}

// Methods (with receivers)
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func (t Triangle) Area() float64 {
	return (t.Base * t.Height) / 2.0
}

// Functions (with parameters)
func Perimeter(width, height float64) float64 {
	return 2 * (width + height)
}
