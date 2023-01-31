package main

import "fmt"

// The (int, int) in this function signature
// shows that the function returns 2 ints.
func vals() (int, int) {
	return 3, 7
}

func main() {
	// Here we use the 2 different return values
    // from the call with multiple assignment.
	a, b := vals()
	fmt.Println("The first returned values is: ", a)
	fmt.Println("The second returned values is: ", b)

	// If you only want a subset of the returned values
    // use the blank identifier _.
	_, c := vals()
	fmt.Println(c)
	
} 