package main

import (
	"fmt"
	"sync"
)

func main() {
	sum := 0
	sumch := make(chan int)
	var wg sync.WaitGroup
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {

			wg.Done()
			sum = sum + 1
			sumch <- sum
		}()
		sum = <-sumch
	}

	wg.Wait()
	fmt.Println(sum)
}
