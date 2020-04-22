package main

import "fmt"

// RoundRectangle defines graphical primitive for drawing rounded rectangles.
type RoundRectangle struct {
	Rectangle
	RoundRadius int64
}

// NewRoundRectangle inits new Round Rectangle and underlying Rectangle.
func NewRoundRectangle(left, top, right, bottom, round int64) *RoundRectangle {
	return &RoundRectangle{
		*NewRectangle(left, top, right, bottom),
		round,
	}
}

// Say returns round rectangle details in special format. Implements Figure.
func (r RoundRectangle) Say() string {
	return fmt.Sprintf("round rectangle: %s and roundRadius=%d", r.Rectangle.Say(), r.RoundRadius)
}
