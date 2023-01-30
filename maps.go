package main

import "fmt"

func main() {
// create an empty map
m := make(map[string]int)
fmt.Println(m)

// set key/value pairs
m["arcade summer"] = 6
m["wild ones"] = 7
fmt.Println(m)

// get a value of a key
v1 := m["arcade summer"]
fmt.Println("el valor entregado con la clave 'arcade summer es: ", v1)

// The builtin len returns the number
// of key/value pairs when called on a map.
fmt.Println(m)

// You can also declare and initialize a new map
// in the same line with this syntax.
n := map[string]int {"foo": 1, "bar": 2}
fmt.Println(n)
	
} 