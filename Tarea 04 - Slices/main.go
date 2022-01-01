package main

import "fmt"

func main() {
	var n int
	var result int
	
	fmt.Scan(&n)

	s := make([]int, n)
	
	for i := 0; i < n; i++{
		fmt.Scan(&s[i])
	}

	for _, v := range s {
		result += v
	}
	fmt.Println("-----------------")
	fmt.Println(result)
}