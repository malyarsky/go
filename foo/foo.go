package main

import (
	"fmt"
	"time"
)

func main() {

	hosts := [...]string{"ahost", "bhost", "chost"}
	size := len(hosts)

	ch := make(chan int, size)
	out := make(chan int, size)

	start := time.Now()
	for i := range hosts {
		ch <- i
		go func() {
			fmt.Println(hosts[<-ch])
			out <- 0
		}()
	}
	for cnt := 0; cnt < size; cnt++ {
		<-out
	}
	fmt.Println(time.Since(start))
}
