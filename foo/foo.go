package main

import (
	"fmt"
	"strings"
)

func main() {
	f := func(r rune) rune {
		return r + 1
	}

	fmt.Println("Incremented runes of a string 'ab': ", strings.Map(f, "ab"))
}
