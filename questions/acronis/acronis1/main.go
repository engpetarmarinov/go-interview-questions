package main

import "fmt"

//FizzBuzz Implementation
//Problem: Write a function that prints the numbers from 1 to 100. But for multiples of three, print "Fizz" instead of the number and for multiples of five print "Buzz". For numbers that are multiples of both three and five, print "FizzBuzz".

func main() {
	FizzBuzz()
	FizzBuzzNumber(5)
}

func FizzBuzzNumber(i int) string {
	if i%3 == 0 && i%5 == 0 {
		return "FizzBuzz"
	}

	if i%3 == 0 {
		return "Fizz"
	}

	if i%5 == 0 {
		return "Buzz"
	}

	return "none"
}

func FizzBuzz() {
	for i := 1; i <= 100; i++ {
		if i%3 == 0 && i%5 == 0 {
			fmt.Println("FizzBuzz")
			continue
		}

		if i%3 == 0 {
			fmt.Println("Fizz")
			continue
		}

		if i%5 == 0 {
			fmt.Println("Buzz")
			continue
		}

		fmt.Println(i)
	}
}
