package main

import "fmt"

func main() {
	var s = make([]int, 0, 10)
	s = append(s, 1, 2, 3, 4, 5)
	fmt.Println(len(s), cap(s))
}