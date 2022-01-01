package main

import "fmt"

var nombre string = "Alex"
var global int64

func main() {
	fmt.Println(nombre)
	fmt.Println(global)
	const apellido string = "Duran"
	fmt.Println(apellido)
	// apellido = "Dominguez"
}
