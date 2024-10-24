package main

import "fmt"

func main() {
	c := gen(1, 2, 3)
	out := sq(c)

	for v := range out {
		fmt.Println(v)
	}
}

func sq(c <-chan int) <-chan int {
	out := make(chan int)

	go func() {
		for n := range c {
			out <- n * n
		}
		close(out)
	}()

	return out
}

func gen(nums ...int) <-chan int {
	out := make(chan int)

	go func() {
		for _, n := range nums {
			out <- n
		}
		close(out)
	}()

	return out
}
