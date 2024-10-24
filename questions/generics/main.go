package main

import "fmt"

func main() {
	fmt.Println(Test[float64](2.0))
}

func Test[T int | float64](t T) T {
	return t + t
}
