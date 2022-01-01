package main

import "fmt"

func main() {
	var e float64
	var n int
	var factorial int

	fmt.Print("Valor de n: ")
	fmt.Scan(&n)

	for i := 0; i <= n; i++ {
		for j := i; j > 1; j-- {
			if j == i {
				factorial = j * (j - 1)
			} else {
				factorial = factorial * (j - 1)
			}
		}
		if i == 0 {
			factorial = 1
		}
		e += (1 / float64(factorial))
	}

	fmt.Println("e = ", e)
}
