package main

import "fmt"

func main() {
	var temp int

	fmt.Println("Temperatura: ")
	fmt.Scan(&temp)

	if temp < 0 {
		fmt.Println("Está helado")
	} else if temp >= 0 && temp < 12 {
		fmt.Println("Está frío")
	} else if temp >= 12 && temp < 18 {
		fmt.Println("Está normal")
	} else{
		fmt.Println("Está caluroso")
	}
}