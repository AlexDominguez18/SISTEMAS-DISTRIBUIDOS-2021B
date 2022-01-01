package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var linea string
	// Instancia de Scanner, regresa un puntero
	scanner := bufio.NewScanner(os.Stdin)
	
	fmt.Print("Nombre de usuario:")
	scanner.Scan()
	
	linea = scanner.Text()

	fmt.Println("Hola", linea)

}