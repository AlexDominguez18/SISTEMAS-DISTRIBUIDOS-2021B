package main

import "fmt"

func getChar() {
	fmt.Print("Presione ENTER para terminar...")
	fmt.Scanln()
}

func main() {
	const PI float32 = 3.1416
	var radius float32
	var area float32

	fmt.Println("CALCULAR AREA DE UN CIRCULO")
	fmt.Print("Radio: ")
	fmt.Scanf("%f", &radius)

	area = PI * (radius * radius)

	fmt.Println("Area = ", area)
	getChar()
}