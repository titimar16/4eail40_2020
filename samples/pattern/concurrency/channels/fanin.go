package main

func fanIn(in1, in2 <-chan int) <-chan int {
	c := make(chan int)
	go func() {
		for {
			c <- <-in1
		}
	}()
	go func() {
		for {
			c <- <-in2
		}
	}()
	return c
}
