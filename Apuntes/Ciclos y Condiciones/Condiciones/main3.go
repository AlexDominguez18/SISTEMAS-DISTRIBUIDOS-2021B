package main

import "fmt"

func main() {
	var opcion int

	fmt.Print("Opcion: ")
	fmt.Scanf("%d", &opcion)

	switch opcion {
	case 1:
		fmt.Println("Opcion 1")
	case 2:
		fmt.Println("Opcion 2")
	case 3:
		fmt.Println("Opcion 3")
	case 4:
		fmt.Println("Opcion 4")
	default:
		fmt.Println("Opcion no valida")
	}

}