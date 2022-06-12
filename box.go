package golang_united_school_homework

import (
	"errors"
)

// box contains list of shapes and able to perform operations on them
type box struct {
	shapes         []Shape
	shapesCapacity int // Maximum quantity of shapes that can be inside the box.
}

// NewBox creates new instance of box
func NewBox(shapesCapacity int) *box {
	return &box{
		shapesCapacity: shapesCapacity,
	}
}

// AddShape adds shape to the box
// returns the error in case it goes out of the shapesCapacity range.
func (b *box) AddShape(shape Shape) error {
	if len(b.shapes) == b.shapesCapacity {
		return errors.New("Added more than max capacity of box")
	}
	b.shapes = append(b.shapes, shape)
	return nil
}

// GetByIndex allows getting shape by index
// whether shape by index doesn't exist or index went out of the range, then it returns an error
func (b *box) GetByIndex(i int) (Shape, error) {
	var shape Shape

	if len(b.shapes)-1 < i {
		return shape, errors.New("Index out of range")
	}

	for idx, val := range b.shapes {
		if i == idx {
			shape = val
		}
	}

	return shape, nil
}

// ExtractByIndex allows getting shape by index and removes this shape from the list.
// whether shape by index doesn't exist or index went out of the range, then it returns an error
func (b *box) ExtractByIndex(i int) (Shape, error) {
	var shape Shape

	if len(b.shapes)-1 < i {
		return shape, errors.New("Index out of range")
	}

	for idx, val := range b.shapes {
		if i == idx {
			shape = val
		}
	}

	b.shapes = append(b.shapes[:i], b.shapes[i+1:]...)

	return shape, nil
}

// ReplaceByIndex allows replacing shape by index and returns removed shape.
// whether shape by index doesn't exist or index went out of the range, then it returns an error
func (b *box) ReplaceByIndex(i int, shape Shape) (Shape, error) {
	var removed Shape
	if len(b.shapes) <= i {
		return removed, errors.New("Index out of range")
	}

	removed = b.shapes[i]
	b.shapes[i] = shape

	return removed, nil
}

// SumPerimeter provides sum perimeter of all shapes in the list.
func (b *box) SumPerimeter() float64 {
	var res float64

	for _, val := range b.shapes {
		res += val.CalcPerimeter()
	}

	return res
}

// SumArea provides sum area of all shapes in the list.
func (b *box) SumArea() float64 {
	var res float64

	for _, val := range b.shapes {
		res += val.CalcArea()
	}

	return res

}

// RemoveAllCircles removes all circles in the list
// whether circles are not exist in the list, then returns an error
func (b *box) RemoveAllCircles() error {

	calCircles := 0

	for i := 0; i < len(b.shapes); i++ {
		switch b.shapes[i].(type) {
		case *Circle:
			calCircles += 1
			b.shapes = append(b.shapes[:i], b.shapes[i+1:]...)
			i -= 1
		}
	}

	if calCircles == 0 {
		return errors.New("No circle in box")
	}

	return nil
}
