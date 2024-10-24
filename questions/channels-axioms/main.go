package main

import (
	"fmt"
	"time"
)

func main() {
	var ch chan int

	go func() {
		//AXIOM 1: sending to nil channel blocks forever
		ch <- 1
		fmt.Println("sent to a nil ch")
	}()

	go func() {
		//AXIOM 2: receiving from nil channel blocks forever
		<-ch
		fmt.Println("received from a nil ch")
	}()

	go func() {
		defer func() {
			if r := recover(); r != nil {
				fmt.Println("recovered from", r)
			}
		}()

		closedChan := make(chan int)
		close(closedChan)
		//AXIOM 3: Pushing to a closed channel will panic
		closedChan <- 1

	}()

	anotherClosedChan := make(chan int)
	close(anotherClosedChan)
	//AXIOM 4: Receiving from closed channel return zero value of the type
	v := <-anotherClosedChan
	fmt.Println("Value from closed chan: ", v)
	time.Sleep(1 * time.Second)
}
