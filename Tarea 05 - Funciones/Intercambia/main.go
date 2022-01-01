package main

import "fmt"

func swap(a* int, b* int) {
	var aux int = *a
	*a = *b
	*b = aux
}

func main() {
	var a int
	var b int

	fmt.Print("a = ")
	fmt.Scan(&a)
	fmt.Print("b = ")
	fmt.Scan(&b)

	fmt.Println("a = ", a, " b = ", b)

	swap(&a, &b)

	fmt.Println("a = ", a, " b = ", b)
}