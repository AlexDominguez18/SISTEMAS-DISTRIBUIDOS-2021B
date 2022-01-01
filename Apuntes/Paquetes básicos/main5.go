package main

import (
	"fmt"
	"sort"
)

func main() {
	s := []int{8, 4, 2, 1, 5}
	sort.Ints(s)
	fmt.Println(s)

	s2 := []int{1, 2, 3, 4}
	sort.Sort(sort.Reverse(sort.IntSlice(s2)))
	fmt.Println(s2)
}