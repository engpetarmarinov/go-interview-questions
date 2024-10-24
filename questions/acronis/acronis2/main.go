//Problem: There is a bug in the following code. Can you spot and fix it?

package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	nums := []int{1, 2, 3, 4, 5}
	wg := &sync.WaitGroup{}

	for i := range nums {
		wg.Add(1)
		go func() {
			defer wg.Done()
			time.Sleep(1 * time.Second)
			fmt.Println(nums[i])
		}()
	}

	wg.Wait()
}
