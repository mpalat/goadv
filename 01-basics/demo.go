package main

import "fmt"

func main() {

	fadd := add
	fpr := func(f func(int, int) int, x, y int) {
		fmt.Println(f(x, y))
	}
	fpr(fadd, 10, 20)

	incr := inc()
	fmt.Println(incr())
	fmt.Println(incr())
	fmt.Println(incr())
	fmt.Println(incr())

}

func add(x, y int) int { return x + y }

func inc() func() int {
	var count = 0
	return func() int {
		count++
		return count
	}
}
