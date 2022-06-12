package golang_united_school_homework

// Rectangle must satisfy to Shape interface
type Rectangle struct {
	Height, Weight float64
}

func (rect Rectangle) CalcArea() float64 {
	area := rect.Height * rect.Weight
	return area
}

func (rect Rectangle) CalcPerimeter() float64 {
	perimeter := rect.Height * 2 + rect.Weight *2
	return perimeter
}
