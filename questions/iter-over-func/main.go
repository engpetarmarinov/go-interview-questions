package main

import (
	"fmt"
	"slices"
)

func MapIncrementBy[T any](in []T, by T) func(func(T) bool) {
	return func(yield func(T) bool) {
		for i, element := range in {
			switch v := any(element).(type) {
			case int:
				valueOfBy := any(by).(int)
				in[i] = any(v + valueOfBy).(T)
				break
			case float64:
				valueOfBy := any(by).(float64)
				in[i] = any(v + valueOfBy).(T)
				break
			default:
			}

			if !yield(in[i]) {
				return
			}
		}
	}
}

func MapPrefix[K comparable, V any](in map[K]V, prefix string) func(func(K, V) bool) {
	return func(yield func(K, V) bool) {
		for key, value := range in {
			switch v := any(value).(type) {
			case string:
				in[key] = any(prefix + v).(V)
				break
			}
			if !yield(key, in[key]) {
				return
			}
		}
	}
}

func main() {
	//iterate over slice
	ints := []int{1, 2, 3, 4, 5}
	fmt.Println("ints")
	for element := range MapIncrementBy(ints, 2) {
		fmt.Println(element)
	}

	floats := []float64{1.1, 2.2, 3.3, 4.4, 5.5}
	fmt.Println("floats")
	for element := range MapIncrementBy(floats, 2.0) {
		fmt.Println(element)
	}

	unordered := []int{6, 2, 5, 1}
	fmt.Println("unordered:", unordered)
	ordered := slices.Sorted(MapIncrementBy(unordered, 1))
	fmt.Println("ordered:", ordered)

	names := map[int]string{
		0: "Vladislavov",
		1: "Marinov",
	}

	//iterate over map
	for _, element := range MapPrefix(names, "Mr.") {
		fmt.Println(element)
	}
}
