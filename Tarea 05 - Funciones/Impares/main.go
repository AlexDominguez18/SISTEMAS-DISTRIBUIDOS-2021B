package main

import "fmt"

func oddNumber() func() uint {
	i := uint(1)

	return func() uint {
		var number = i
		i += 2
		return number
	}
}

func main() {
	var n int

	fmt.Scan(&n)
	nextNumber := oddNumber()

	for i := 0; i < n; i++ {
		fmt.Print(nextNumber(), " ")
	}

	fmt.Println()
}