package main

import "fmt"

func main() {
	var numero float64

	fmt.Println("Float: ")
	// fmt.Scanf("%f", &numero)
	fmt.Scan(&numero)

	output := numero * 2

	fmt.Println(output)
}