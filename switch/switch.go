package main

import (
	"fmt"
)

func main() {
	var Foo = func(n int) int {
		fmt.Println(n)
		return n
	}
	switch 4 {
	case Foo(1), Foo(2), Foo(3):
		fmt.Println("Первый case")
		//fallthrough
	case Foo(4):
		fmt.Println("Второй case")
	}
}
