package circle

import (
	"fmt"
	"math"
)

// Circle defines graphical primitive for drawing circles.
type Circle struct {
	X, Y, Radius int64
}

// Init - an empty initialization
func Init() {}

// NewCircle inits new Circle.
func NewCircle(x, y, radius int64) *Circle {
	return &Circle{
		X:      x,
		Y:      y,
		Radius: radius,
	}
}

// Say returns circle details in special format. Implements Figure.
func (c Circle) Say() string {
	return fmt.Sprintf("circle: radius=%d and centre=(%d,%d)", c.Radius, c.X, c.Y)
}

// Square returns square of the circle. Implements Figure.
func (c Circle) Square() float64 {
	return math.Pi * math.Pow(float64(c.Radius), 2)
}
