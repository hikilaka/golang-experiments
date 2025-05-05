package fib

import (
	"fmt"
	"time"
)

func fibonacci(n uint) uint {
	if n == 0 || n == 1 {
		return n
	}

	return fib(n-2) + fib(n-1)
}

func memoize(fn func(uint) uint) func(uint) uint {
	cache := []uint{}

	return func(i uint) uint {
		if len(cache) >= int(i) {
			return cache[i]
		}

		result := fn(i)
		cache = append(cache, result)

		return result
	}
}

func timeFuncCall(name string, fn func(uint) uint, param uint) {
	start := time.Now()

	result := fn(param)

	elapsed := time.Since(start)
	fmt.Printf("%s(%d)=%d - took %s\n", name, param, result, elapsed)
}

func main() {
	timeFuncCall("recursive_fib", fibonacci, 15)
	timeFuncCall("recursive_fib", fibonacci, 25)
	timeFuncCall("recursive_fib", fibonacci, 45)

	f := memoize(fibonacci)
	timeFuncCall("memoized_fib", f, 15)
	timeFuncCall("memoized_fib", f, 25)
	timeFuncCall("memoized_fib", f, 45)
}
