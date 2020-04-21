package main

import "fmt"

// Figure describes graphical primitive, which can Say
// own information and return it's Square.
type Figure interface {
	Say() string
	Square() float64
}

func main() {
	figures := []Figure{
		NewRectangle(10, 20, 600, 400),
		NewCircle(50, 300, 150),
		NewRoundRectangle(30, 40, 500, 200, 5),
	}
	for _, figure := range figures {
		fmt.Println(figure.Say())
	}
}
