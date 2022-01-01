package main

import "fmt"

func getChar() {
	fmt.Print("Presione ENTER para terminar...")
	fmt.Scanln()
}

func main() {
	var farenheitDegrees float32
	var celciusDegrees float32

	fmt.Println("FAHRENHEIT A CELCIUS")
	fmt.Print("F°: ")
	fmt.Scanf("%f", &farenheitDegrees)

	celciusDegrees = (farenheitDegrees - 32) * 5 / 9

	fmt.Printf("%.2f F° equivalen a %.2f C°\n", farenheitDegrees, celciusDegrees)
	getChar()
}