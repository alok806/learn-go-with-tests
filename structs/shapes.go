package main

import "math"

/*
Rectangle struct
*/
type Rectangle struct {
	Width  float64
	Height float64
}

/*
Area of rectangle
*/
func (r Rectangle) Area() float64 {
	return r.Height * r.Width
}

/*
Circle struct
*/
type Circle struct {
	Radius float64
}

/*
Area of circle
*/
func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

/*
Perimeter of rectangle
*/
func Perimeter(rectangle Rectangle) float64 {
	return 2 * (rectangle.Width + rectangle.Height)
}

/*
Shape interface for area of a shape
*/
type Shape interface {
	Area() float64
}

/*
Triangle struct
*/
type Triangle struct {
	Base   float64
	Height float64
}

/*
Area of a triangle
*/
func (t Triangle) Area() float64 {
	return t.Base * t.Height * 0.5
}
