package main

import (
	"fmt"
)

// reverse reverses a slice order
func reverse(s []int) []int {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}

	return s
}

func main() {
	s := []int{1, 2, 3, 4, 5}
	fmt.Println(reverse(s))
}
