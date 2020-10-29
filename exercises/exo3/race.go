package main

import (
	"fmt"
	"sync"
)

// Fix the race condition in the code

var counter int

func main() {
	var wg sync.WaitGroup

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go countUp(&wg)
	}
	wg.Wait()
	fmt.Printf("counting %d", counter)
}

func countUp(wg *sync.WaitGroup) {
	defer wg.Done()
	counter = counter + 1
}
