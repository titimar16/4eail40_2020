package main

import "fmt"

func main() {
	a := adder()

	for i := 0; i < 10; i++ {
		fmt.Println(a())
	}
}

func adder() func() int {
	var counter int
	return func() int {
		counter++
		return counter
	}
}
