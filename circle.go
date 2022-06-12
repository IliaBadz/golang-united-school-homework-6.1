package golang_united_school_homework

import "math"

// Circle must satisfy to Shape interface
type Circle struct {
	Radius float64
}

func (circle Circle) CalcArea() float64 {
	area := math.Pow(circle.Radius, 2) * math.Pi
	return area
}

func (circle Circle) CalcPerimeter() float64 {
	perimeter := 2 * math.Pi * circle.Radius
	return perimeter
}
