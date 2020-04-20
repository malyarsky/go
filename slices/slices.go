package main

import (
	"fmt"
)

func main() {
	a := [...]int{0, 1, 2, 3}
	b := [...]int{0, 1, 2, 3}
	// !works for arrays but slices
	fmt.Println(a == b)

	m := [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(m)
	s1 := m[1:8]
	fmt.Println(s1)
	for i := range s1 {
		s1[i] *= 10
	}
	fmt.Println(s1)
	fmt.Println(m)
}
