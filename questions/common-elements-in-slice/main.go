package main

import (
	"fmt"
)

func main() {
	a := []string{"a", "b", "c", "d", "e"}
	b := []string{"c", "d", "e", "f", "g"}

	c := getCommonElementsFromTwoSlices(a, b)

	fmt.Println(c)
}

func getCommonElementsFromTwoSlices(a []string, b []string) []string {
	var result []string
	mapped := make(map[string]interface{})

	for _, v := range a {
		mapped[v] = nil
	}

	for _, v := range b {
		if _, ok := mapped[v]; ok {
			result = append(result, v)
		}
	}

	return result
}
