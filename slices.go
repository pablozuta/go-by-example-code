package main

import "fmt"

func main() {
	 // Unlike arrays, slices are typed only by the elements they contain
    // (not the number of elements). To create an empty slice with non-zero length
    // use the builtin make. Here we make a slice of strings of length 3
    // (initially zero-valued).
	s := make([]string, 3)
	fmt.Println("emp: ", s)

	// We can set and get just like with arrays.
	s[0] = "cyberpunk"
	s[1] = "synthwave"
	s[2] = "darkwave"
	fmt.Println("set: ", s)
	fmt.Println("get: ", s[2])

	// len returns the length of the slice as expected.
	fmt.Println(len(s))

	// In addition to these basic operations, slices support several more that
    // make them richer than arrays. One is the builtin append, which returns
    // a slice containing one or more new values. Note that we need to accept
    // a return value from append as we may get a new slice value.
	s = append(s, "retrowave")
	s = append(s, "atlas", "FM-84")
	fmt.Println(s)

	// Slices can also be copy’d. Here we create an empty slice c of the same length
    // as s and copy into c from s.
	c :=make([]string, len(s))
	copy(c, s)
	fmt.Println("copy of slice s: ", c)

	// Slices support a “slice” operator with the syntax slice[low:high].
    // For example, this gets a slice of the elements s[2], s[3], and s[4].
	l := s[2:5]
	fmt.Println(l)

	// We can declare and initialize a variable for slice in a single line as well.
	t := []string {"everything", "running in the night", "tears", "chasing yesterday"}
	fmt.Println(t)

	// Slices can be composed into multi-dimensional data structures.
    // The length of the inner slices can vary, unlike with multi-dimensional arrays.
	twoD := make([][]int, 3)
    for i := 0; i < 3; i++ {
        innerLen := i + 1
        twoD[i] = make([]int, innerLen)
        for j := 0; j < innerLen; j++ {
            twoD[i][j] = i + j
        }
    }
    fmt.Println("2D slice: ", twoD)

} 