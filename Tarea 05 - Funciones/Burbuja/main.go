package main

import "fmt"

func bubbleSort(s []int64) {
	for i := 0; i < len(s); i++ {
		for j := 0; j < len(s) - i - 1; j++ {
			if s[j] > s[j + 1] {
				aux := s[j]
				s[j] = s[j + 1]
				s[j + 1] = aux
			}
		}
	}
}

func main() {
	s := []int64 {22, 100, 7, 11, 20, 34, 5}

	fmt.Println("S inicial = ", s)

	bubbleSort(s)

	fmt.Println("S ordenado = ", s)
}