package main

import "fmt"

func main() {
	x := [...]int{1, 2, 3, 4} //Arreglo al que podmos seguir añadiendo valores
	// x := [5]int{1,2,3,4,5}
	// x[4] = 100
	fmt.Println(x)
	fmt.Println(len(x), cap(x))

	// for i := 0; i < len(x); i++ {
	// 	fmt.Println(x[i])
	// }

	// for i, v := range x {
	// 	fmt.Println(v, i)
	// }

	// for _, v := range x {
	// 	fmt.Println(v)
	// }

	// for i, _ := range x {
	// 	fmt.Println(i)
	// }
}