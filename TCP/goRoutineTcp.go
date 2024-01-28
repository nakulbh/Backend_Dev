package main

import (
	"fmt"
	"sync"
)

func worker(ports chan int, wg *sync.WaitGroup) {
	for p := range ports {
		fmt.Println(p)
		wg.Done()

	}
}

func main() {
	ports := makw(chan int, 100)
	/* You’ve capped the channel at 100,
	meaning it can hold 100 items before the sender will block.
	*/
	var wg sync.WaitGroup
	for i := 0; i < 1024; i++ {
		wg.Add(1)
		ports <- i

	}
	wg.Wait()
	close(ports)
}
