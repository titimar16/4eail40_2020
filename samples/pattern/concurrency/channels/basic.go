package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan int)
	close(c)
	_, ok := <-c
	fmt.Println(ok) // Print false
}

func tryReceive(c <-chan int) (data int, more, ok bool) {
	select {
	case data, more = <-c:
		return data, more, true
	default:
		return 0, true, false
	}
}

func tryReceiveWithTimeout(c <-chan int, duration time.Duration) (data int, more, ok bool) {
	select {
	case data, more = <-c:
		return data, more, true
	case <-time.After(duration): // returns a channel
		return 0, true, false
	}
}
