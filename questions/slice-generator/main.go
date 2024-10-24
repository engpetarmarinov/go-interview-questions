package main

import "fmt"

func main() {
	for _, i := range sliceInt(10, 20) {
		fmt.Println(i)
	}
}

func sliceInt(from int, to int) []int {
	s := make([]int, 0, to-from)
	for i := from; i < to; i++ {
		s = append(s, i)
	}
	return s
}
