package main

import "fmt"

func fib(n int) int {
	var mf = memoizedIntToInt(fib)
	if n <= 1 {
		return 1
	}
	return mf(n-2) + mf(n-1)
}

var cache = make(map[int]int) // should be unbounded
func memoizedIntToInt(f func(int) int) func(int) int {
	return func(i int) int {
		if cached, exist := cache[i]; exist {
			return cached
		}
		res := f(i)
		cache[i] = res
		return res
	}
}

func main() {
	fmt.Println(fib(45))

}
