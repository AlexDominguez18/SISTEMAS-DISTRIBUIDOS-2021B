package main

import "fmt"

func getChar() {
	fmt.Print("Presiona ENTER para terminar...")
	fmt.Scanln()
}

func main() {
	var width float32
	var height float32
	var area float32

	fmt.Println("CALCULAR EL AREA DE UN TRIANGULO")
	fmt.Print("Base: ")
	fmt.Scanf("%f", &width)
	fmt.Print("Altura: ")
	fmt.Scanf("%f", &height)
	
	area = (width * height) / 2

	fmt.Println("Area = ", area)
	getChar()
}