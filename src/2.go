package main

import "fmt"

func main() {
	// A function that takes two numbers as input and returns their sum
	sum := func(a int, b int) int {
		return a + b
	}

	// Test the function with different inputs
	fmt.Println("1 + 2 =", sum(1, 2))
	fmt.Println("3 + 4 =", sum(3, 4))
	fmt.Println("5 + 6 =", sum(5, 6))
}
