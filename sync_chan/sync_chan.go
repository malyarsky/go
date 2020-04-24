package main

import (
	"fmt"
	"sync"
	"time"
)

// Достигается параллельность выполнения всех рутин и при этом требуемая последовательность вывода без потери параллельности
func main() {

	hosts := [...]string{"ahost", "bhost", "chost"}
	size := len(hosts)

	start := time.Now()

	ch := make(chan int, size)
	//	out := make(chan int, size)
	var wg sync.WaitGroup

	for i := range hosts {
		ch <- i
		wg.Add(1)
		go func() {
			fmt.Println(hosts[<-ch])
			wg.Done()
			//out <- 0
		}()
	}
	/*
		for cnt := 0; cnt < size; cnt++ {
			<-out
		}
	*/
	wg.Wait()
	fmt.Println(time.Since(start))
}
