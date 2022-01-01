package main

import "fmt"

func getChar() {
	fmt.Print("Presione ENTER para terminar...")
	fmt.Scanln()
}

func main() {
	var width float32
	var area float32
	
	fmt.Println("CALCULAR AREA DE UN CUADRADO")
	fmt.Print("Base: ")
	fmt.Scanf("%f", &width)

	area = width * width

	fmt.Println("Area = ", area)
	getChar()
}