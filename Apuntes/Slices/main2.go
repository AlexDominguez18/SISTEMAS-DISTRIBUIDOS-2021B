package main

import "fmt"

// SLICES

func main() {
	// x := [...]int{1, 2, 3, 4, 5, 6, 7}
	// s := x[0:4] // Slice
	// fmt.Println(len(s), cap(s))
	var s []int
	fmt.Println(len(s), cap(s))
	s = append(s, 1, 2, 3, 4, 5, 6)

	fmt.Println(len(s), cap(s))
	for _, v := range s {
		fmt.Println(v)
	}
}