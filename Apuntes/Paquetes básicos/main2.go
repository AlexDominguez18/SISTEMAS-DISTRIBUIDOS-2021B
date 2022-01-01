package main

import (
	"fmt"
	"os"
)

func main() {
	//Instancia de archivo o error
	file, err := os.Create("test.txt")

	if err != nil {
		fmt.Print(err)
		return
	}

	defer file.Close()

	file.WriteString("Hola mundo!")
}