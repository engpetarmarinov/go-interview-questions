package main

import (
	"fmt"
)

func main() {
	fmt.Println(majorityElementVer6([]int{1, 1, 3, 3, 3, 3, 3, 3, 3, 2, 3, 4, 5, 6, 6, 3, 3, 3, 3, 3}))
}

/*
Given an array nums of size n, return the majority element.

The majority element is the element that appears more than ⌊n / 2⌋ times. You may assume that the majority element always exists in the array.
*/
func majorityElementVer6(nums []int) int {
	candidate := 0
	count := 0

	for _, num := range nums {
		if count == 0 {
			candidate = num

			count = 0
			for _, num := range nums {
				if num == candidate {
					count++
				}
			}

			if count > len(nums)/2 {
				return candidate
			}

			return -1
		}
	}

	return 0
}
