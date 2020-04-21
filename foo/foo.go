package main

import "fmt"

// Person is
type Person struct {
	Name string
	Age  int
}

type 

func (p Person) String() string {
	return fmt.Sprintf("%v (%v года)", p.Name, p.Age)
}

func main() {
	a := Person{"Иван Иванов", 54}
	z := &Person{"Петр Петров", 2004}
	fmt.Println(a, z)
	fmt.Println(&a, z)
}
