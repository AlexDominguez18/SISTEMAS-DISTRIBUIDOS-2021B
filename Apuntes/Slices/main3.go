package main

import "fmt"

func main() {
	 s := make([]int, 5, 100) //Reserva dinamica de memoria, similar al new en C++
	 fmt.Println(len(s), cap(s))
	 s = append(s, 1, 2, 3, 4, 5) //Reserva dinamica
	 fmt.Println(len(s), cap(s))
}