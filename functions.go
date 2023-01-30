package main

import "fmt"

// Here’s a function that takes two ints and returns their sum
func plus(a int, b int) int {
	return a + b
}

func multiplier(a int, b int) int {
	return a * b
}

// When you have multiple consecutive parameters of the same type
// you may omit the type name for the like-typed parameters
func plusPlus (a, b, c int) int {
	return a + b + c
}
func main() {
	
	fmt.Println(plus(2, 76))
	fmt.Println(plusPlus(2, 76, 32))

	fmt.Println("23 * 34 es: ", multiplier(23, 34))
	
} 