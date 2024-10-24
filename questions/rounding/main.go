package main

import "fmt"

func main() {
	var a float32 = 1.0000001
	var b float32 = 1.0000002
	var c float32 = b - a

	fmt.Println("a:", a)        // 1.0000001
	fmt.Println("b:", b)        // 1.0000002
	fmt.Printf("c: %.10f\n", c) // 0.0000001192
}
