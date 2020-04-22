package rectangle

import (
	"fmt"
	"math"
)

// Rectangle defines graphical primitive for drawing rectangles.
type Rectangle struct {
	Left, Right, Top, Bottom int64
}

// NewRectangle inits new Rectangle.
func NewRectangle(left, top, right, bottom int64) *Rectangle {
	return &Rectangle{
		Left:   left,
		Top:    top,
		Right:  right,
		Bottom: bottom,
	}
}

// Say returns rectangle details in special format. Implements Figure.
func (r Rectangle) Say() string {
	return fmt.Sprintf("rectangle: Rect (left=%d,top=%d,right=%d,bottom=%d)", r.Left, r.Top, r.Right, r.Bottom)
}

// Square returns square of the rectangle. Implements Figure.
func (r Rectangle) Square() float64 {
	return math.Abs(float64((r.Right - r.Left) * (r.Top - r.Bottom)))
}
