package main

/*
Rectangle struct
*/
type Rectangle struct {
	Width  float64
	Height float64
}

/*
Perimeter of rectangle
*/
func Perimeter(rectangle Rectangle) float64 {
	return 2 * (rectangle.Width + rectangle.Height)
}

/*
Area of rectangle
*/
func Area(rectangle Rectangle) float64 {
	return rectangle.Width * rectangle.Height
}
