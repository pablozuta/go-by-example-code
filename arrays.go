package main

import "fmt"

func main() {
	// Here we create an array a that will hold exactly 5 ints.
	var a [5]int
	fmt.Println("emp: ", a)
	a[4] = 100
	fmt.Println(a)

	// The builtin len returns the length of an array.
	fmt.Println("len: ", len(a))

	// declare and initialize
	b := [5]int {2, 6, 544, 65}
	fmt.Println(b)

	
	
} 