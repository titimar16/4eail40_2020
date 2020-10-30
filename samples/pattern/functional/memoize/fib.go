package memoize

func fibonacci(n int) int {
	if n <= 1 {
		return 1
	}
	return memoizedIntInt(fibonacci)(n-1) + memoizedIntInt(fibonacci)(n-2)
}

var cache = make(map[int]int)

func memoizedIntInt(f func(int) int) func(int) int {
	return func(i int) int {
		if cached, exist := cache[i]; exist {
			return cached
		}
		ret := f(i)
		cache[ret] = ret
		return ret
	}
}
