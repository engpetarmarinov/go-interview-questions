package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	// Allocate a large amount of memory
	for i := 0; i < 10; i++ {
		s := make([]byte, 10*1024*1024) // 10 MB
		_ = s
	}

	var memStats runtime.MemStats
	runtime.ReadMemStats(&memStats)

	fmt.Println(memStats.Alloc)

	// Trigger a garbage collection cycle manually
	//fmt.Println("Forcing garbage collection")
	//runtime.GC()

	go func() {
		tick := time.Tick(time.Second)
		for {
			select {
			case <-tick:
				fmt.Println(time.Now().String(), memStats.Alloc)
			}
		}
	}()

	// Wait to observe GC behavior
	time.Sleep(2222 * time.Second)
}
