package golang_united_school_homework

import "math"

// Triangle must satisfy to Shape interface
type Triangle struct {
	Side float64
}

func (triangle Triangle) CalcArea() float64 {
	area := math.Pow(triangle.Side, 2) * math.Sqrt(3) / 4
	return area
}

func (triangle Triangle) CalcPerimeter() float64 {
	perimeter := triangle.Side * 3
	return perimeter
}
