package main

import "fmt"

func main() {
	arr := []int{1, 3, 5, 7, 9, 11}
	index := binarySearch(arr, 12)

	if index == -1 {
		fmt.Println("not found")
	}

	fmt.Println(index)
}

func binarySearch(arr []int, i int) int {
	low, high := 0, len(arr)-1

	for low <= high {
		mid := low + (high-low)/2
		if arr[mid] == i {
			return mid
		}

		if arr[mid] > i {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}

	return -1
}
