package structs

import "math"

// A struct is a named collection of fields to store data
type Rectangle struct {
	Width float64
	Height float64
}

type Circle struct {
	Radius float64
}

type Triangle struct {
	Base float64
	Height float64
}

// A new type (interface rather than struct)
type Shape interface {
	Area() float64
}

// Method is a function with a receiver:
// func (receiverName ReceiverType) MethodName(args)
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func (t Triangle) Area() float64 {
	return (t.Base * t.Height) * 0.5
}

func Perimeter(rectangle Rectangle) float64 {
	return 2 * (rectangle.Width + rectangle.Height)
}
